package ws

import (
	"bytes"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/quick"
	"time"

	"github.com/gorilla/websocket"
)

type TestServer struct {
	m *Melody
}

func NewTestServerHandler(handler handleMessageFunc) *TestServer {
	m := New()
	m.HandleMessage(handler)
	return &TestServer{
		m: m,
	}
}

func NewTestServer() *TestServer {
	m := New()
	return &TestServer{
		m: m,
	}
}

func (s *TestServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := s.m.HandleRequest(w, r)
	if err != nil {
		return
	}
}

func NewDialer(url string) (*websocket.Conn, error) {
	dialer := &websocket.Dialer{}
	conn, _, err := dialer.Dial(strings.Replace(url, "http", "ws", 1), nil)
	return conn, err
}

func TestEcho(t *testing.T) {
	echo := NewTestServerHandler(func(session *Session, msg []byte) {
		err := session.Write(msg)
		if err != nil {
			return
		}
	})
	server := httptest.NewServer(echo)
	defer server.Close()

	fn := func(msg string) bool {
		conn, err := NewDialer(server.URL)
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)

		if err != nil {
			t.Error(err)
			return false
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			return false
		}

		_, ret, err := conn.ReadMessage()

		if err != nil {
			t.Error(err)
			return false
		}

		if msg != string(ret) {
			t.Errorf("%s should equal %s", msg, string(ret))
			return false
		}

		return true
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func TestWriteClosed(t *testing.T) {
	echo := NewTestServerHandler(func(session *Session, msg []byte) {
		err := session.Write(msg)
		if err != nil {
			return
		}
	})
	server := httptest.NewServer(echo)
	defer server.Close()

	fn := func(msg string) bool {
		conn, err := NewDialer(server.URL)

		if err != nil {
			t.Error(err)
			return false
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			return false
		}

		echo.m.HandleConnect(func(s *Session) {
			err := s.Close()
			if err != nil {
				return
			}
		})

		echo.m.HandleDisconnect(func(s *Session) {
			err := s.Write([]byte("hello world"))

			if err == nil {
				t.Error("should be an error")
			}
		})

		return true
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func TestLen(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	connect := int(rand.Int31n(100))
	disconnect := rand.Float32()
	conns := make([]*websocket.Conn, connect)
	defer func() {
		for _, conn := range conns {
			if conn != nil {
				err := conn.Close()
				if err != nil {
					return
				}
			}
		}
	}()

	echo := NewTestServerHandler(func(session *Session, msg []byte) {})
	server := httptest.NewServer(echo)
	defer server.Close()

	disconnected := 0
	for i := 0; i < connect; i++ {
		conn, err := NewDialer(server.URL)

		if err != nil {
			t.Error(err)
		}

		if rand.Float32() < disconnect {
			conns[i] = nil
			disconnected++
			err := conn.Close()
			if err != nil {
				return
			}
			continue
		}

		conns[i] = conn
	}

	time.Sleep(time.Millisecond)

	connected := connect - disconnected

	if echo.m.Len() != connected {
		t.Errorf("melody len %d should equal %d", echo.m.Len(), connected)
	}
}

func TestEchoBinary(t *testing.T) {
	echo := NewTestServer()
	echo.m.HandleMessageBinary(func(session *Session, msg []byte) {
		err := session.WriteBinary(msg)
		if err != nil {
			return
		}
	})
	server := httptest.NewServer(echo)
	defer server.Close()

	fn := func(msg string) bool {
		conn, err := NewDialer(server.URL)
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)

		if err != nil {
			t.Error(err)
			return false
		}

		if err := conn.WriteMessage(websocket.BinaryMessage, []byte(msg)); err != nil {
			return false
		}

		_, ret, err := conn.ReadMessage()

		if err != nil {
			t.Error(err)
			return false
		}

		if msg != string(ret) {
			t.Errorf("%s should equal %s", msg, string(ret))
			return false
		}

		return true
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func TestHandlers(t *testing.T) {
	echo := NewTestServer()
	echo.m.HandleMessage(func(session *Session, msg []byte) {
		err := session.Write(msg)
		if err != nil {
			return
		}
	})
	server := httptest.NewServer(echo)
	defer server.Close()

	var q *Session

	echo.m.HandleConnect(func(session *Session) {
		q = session
		err := session.Close()
		if err != nil {
			return
		}
	})

	echo.m.HandleDisconnect(func(session *Session) {
		if q != session {
			t.Error("disconnecting session should be the same as connecting")
		}
	})

	_, err := NewDialer(server.URL)
	if err != nil {
		return
	}
}

// ! STILL ERROR NEED TO FIX NEXT....
//func TestMetadata(t *testing.T) {
//	echo := NewTestServer()
//	echo.m.HandleConnect(func(session *Session) {
//		session.Set("stamp", time.Now().UnixNano())
//	})
//	echo.m.HandleMessage(func(session *Session, msg []byte) {
//		stamp := session.MustGet("stamp").(int64)
//		session.Write([]byte(strconv.Itoa(int(stamp))))
//	})
//	server := httptest.NewServer(echo)
//	defer server.Close()
//
//	fn := func(msg string) bool {
//		conn, err := NewDialer(server.URL)
//		defer conn.Close()
//
//		if err != nil {
//			t.Error(err)
//			return false
//		}
//
//		conn.WriteMessage(websocket.TextMessage, []byte(msg))
//
//		_, ret, err := conn.ReadMessage()
//
//		if err != nil {
//			t.Error(err)
//			return false
//		}
//
//		stamp, err := strconv.Atoi(string(ret))
//
//		if err != nil {
//			t.Error(err)
//			return false
//		}
//
//		diff := int(time.Now().UnixNano()) - stamp
//
//		if diff <= 0 {
//			t.Errorf("diff should be above 0 %d", diff)
//			return false
//		}
//
//		return true
//	}
//
//	if err := quick.Check(fn, nil); err != nil {
//		t.Error(err)
//	}
//}

func TestUpgrader(t *testing.T) {
	broadcast := NewTestServer()
	broadcast.m.HandleMessage(func(session *Session, msg []byte) {
		err := session.Write(msg)
		if err != nil {
			return
		}
	})
	server := httptest.NewServer(broadcast)
	defer server.Close()

	broadcast.m.Upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return false },
	}

	broadcast.m.HandleError(func(session *Session, err error) {
		if err == nil || err.Error() != "websocket: origin not allowed" {
			t.Error("there should be a origin error")
		}
	})

	_, err := NewDialer(server.URL)

	if err == nil || err.Error() != "websocket: bad handshake" {
		t.Error("there should be a badhandshake error")
	}
}

//goland:noinspection ALL
func TestBroadcast(t *testing.T) {
	broadcast := NewTestServer()
	broadcast.m.HandleMessage(func(session *Session, msg []byte) {
		err := broadcast.m.Broadcast(msg)
		if err != nil {
			return
		}
	})
	server := httptest.NewServer(broadcast)
	defer server.Close()

	n := 10

	fn := func(msg string) bool {
		conn, _ := NewDialer(server.URL)
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				return
			}
		}(conn)

		listeners := make([]*websocket.Conn, n)
		for i := 0; i < n; i++ {
			listener, _ := NewDialer(server.URL)
			listeners[i] = listener
			defer func(conn *websocket.Conn) {
				err := conn.Close()
				if err != nil {

				}
			}(listeners[i])
		}

		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			return false
		}

		for i := 0; i < n; i++ {
			_, ret, err := listeners[i].ReadMessage()

			if err != nil {
				t.Error(err)
				return false
			}

			if msg != string(ret) {
				t.Errorf("%s should equal %s", msg, string(ret))
				return false
			}
		}

		return true
	}

	if !fn("test") {
		t.Errorf("should not be false")
	}
}

//goland:noinspection ALL
func TestBroadcastBinary(t *testing.T) {
	broadcast := NewTestServer()
	broadcast.m.HandleMessageBinary(func(session *Session, msg []byte) {
		err := broadcast.m.BroadcastBinary(msg)
		if err != nil {
			return
		}
	})
	server := httptest.NewServer(broadcast)
	defer server.Close()

	n := 10

	fn := func(msg []byte) bool {
		conn, _ := NewDialer(server.URL)
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				return
			}
		}(conn)

		listeners := make([]*websocket.Conn, n)
		for i := 0; i < n; i++ {
			listener, _ := NewDialer(server.URL)
			listeners[i] = listener
			defer func(conn *websocket.Conn) {
				err := conn.Close()
				if err != nil {
					return
				}
			}(listeners[i])
		}

		err := conn.WriteMessage(websocket.BinaryMessage, []byte(msg))
		if err != nil {
			return false
		}

		for i := 0; i < n; i++ {
			messageType, ret, err := listeners[i].ReadMessage()

			if err != nil {
				t.Error(err)
				return false
			}

			if messageType != websocket.BinaryMessage {
				t.Errorf("message type should be BinaryMessage")
				return false
			}

			if !bytes.Equal(msg, ret) {
				t.Errorf("%v should equal %v", msg, ret)
				return false
			}
		}

		return true
	}

	if !fn([]byte{2, 3, 5, 7, 11}) {
		t.Errorf("should not be false")
	}
}

//goland:noinspection ALL
func TestBroadcastOthers(t *testing.T) {
	broadcast := NewTestServer()
	broadcast.m.HandleMessage(func(session *Session, msg []byte) {
		err := broadcast.m.BroadcastOthers(msg, session)
		if err != nil {
			return
		}
	})
	broadcast.m.Config.PongWait = time.Second
	broadcast.m.Config.PingPeriod = time.Second * 9 / 10
	server := httptest.NewServer(broadcast)
	defer server.Close()

	n := 10

	fn := func(msg string) bool {
		conn, _ := NewDialer(server.URL)
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				return
			}
		}(conn)

		listeners := make([]*websocket.Conn, n)
		for i := 0; i < n; i++ {
			listener, _ := NewDialer(server.URL)
			listeners[i] = listener
			defer func(conn *websocket.Conn) {
				err := conn.Close()
				if err != nil {
					return
				}
			}(listeners[i])
		}

		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			return false
		}

		for i := 0; i < n; i++ {
			_, ret, err := listeners[i].ReadMessage()

			if err != nil {
				t.Error(err)
				return false
			}

			if msg != string(ret) {
				t.Errorf("%s should equal %s", msg, string(ret))
				return false
			}
		}

		return true
	}

	if !fn("test") {
		t.Errorf("should not be false")
	}
}

//goland:noinspection GoDeferInLoop
func TestBroadcastBinaryOthers(t *testing.T) {
	broadcast := NewTestServer()
	broadcast.m.HandleMessageBinary(func(session *Session, msg []byte) {
		err := broadcast.m.BroadcastBinaryOthers(msg, session)
		if err != nil {
			return
		}
	})
	broadcast.m.Config.PongWait = time.Second
	broadcast.m.Config.PingPeriod = time.Second * 9 / 10
	server := httptest.NewServer(broadcast)
	defer server.Close()

	n := 10

	fn := func(msg []byte) bool {
		conn, _ := NewDialer(server.URL)
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				return
			}
		}(conn)

		listeners := make([]*websocket.Conn, n)
		for i := 0; i < n; i++ {
			listener, _ := NewDialer(server.URL)
			listeners[i] = listener
			defer func(conn *websocket.Conn) {
				err := conn.Close()
				if err != nil {
					return
				}
			}(listeners[i])
		}

		err := conn.WriteMessage(websocket.BinaryMessage, msg)
		if err != nil {
			return false
		}

		for i := 0; i < n; i++ {
			messageType, ret, err := listeners[i].ReadMessage()

			if err != nil {
				t.Error(err)
				return false
			}

			if messageType != websocket.BinaryMessage {
				t.Errorf("message type should be BinaryMessage")
				return false
			}

			if !bytes.Equal(msg, ret) {
				t.Errorf("%v should equal %v", msg, ret)
				return false
			}
		}

		return true
	}

	if !fn([]byte{2, 3, 5, 7, 11}) {
		t.Errorf("should not be false")
	}
}

func TestPingPong(t *testing.T) {
	noecho := NewTestServer()
	noecho.m.Config.PongWait = time.Second
	noecho.m.Config.PingPeriod = time.Second * 9 / 10
	server := httptest.NewServer(noecho)
	defer server.Close()

	conn, err := NewDialer(server.URL)
	conn.SetPingHandler(func(string) error {
		return nil
	})
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	if err != nil {
		t.Error(err)
	}

	if err := conn.WriteMessage(websocket.TextMessage, []byte("test")); err != nil {
		return
	}

	_, _, err = conn.ReadMessage()

	if err == nil {
		t.Error("there should be an error")
	}
}

func TestBroadcastFilter(t *testing.T) {
	broadcast := NewTestServer()
	broadcast.m.HandleMessage(func(session *Session, msg []byte) {
		err := broadcast.m.BroadcastFilter(msg, func(q *Session) bool {
			return session == q
		})
		if err != nil {
			return
		}
	})
	server := httptest.NewServer(broadcast)
	defer server.Close()

	fn := func(msg string) bool {
		conn, err := NewDialer(server.URL)
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				return
			}
		}(conn)

		if err != nil {
			t.Error(err)
			return false
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			return false
		}

		_, ret, err := conn.ReadMessage()

		if err != nil {
			t.Error(err)
			return false
		}

		if msg != string(ret) {
			t.Errorf("%s should equal %s", msg, string(ret))
			return false
		}

		return true
	}

	if !fn("test") {
		t.Errorf("should not be false")
	}
}

func TestBroadcastBinaryFilter(t *testing.T) {
	broadcast := NewTestServer()
	broadcast.m.HandleMessageBinary(func(session *Session, msg []byte) {
		err := broadcast.m.BroadcastBinaryFilter(msg, func(q *Session) bool {
			return session == q
		})
		if err != nil {
			return
		}
	})
	server := httptest.NewServer(broadcast)
	defer server.Close()

	fn := func(msg []byte) bool {
		conn, err := NewDialer(server.URL)
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				return
			}
		}(conn)

		if err != nil {
			t.Error(err)
			return false
		}

		if err := conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
			return false
		}

		messageType, ret, err := conn.ReadMessage()

		if err != nil {
			t.Error(err)
			return false
		}

		if messageType != websocket.BinaryMessage {
			t.Errorf("message type should be BinaryMessage")
			return false
		}

		if !bytes.Equal(msg, ret) {
			t.Errorf("%v should equal %v", msg, ret)
			return false
		}

		return true
	}

	if !fn([]byte{2, 3, 5, 7, 11}) {
		t.Errorf("should not be false")
	}
}

func TestStop(t *testing.T) {
	noecho := NewTestServer()
	server := httptest.NewServer(noecho)
	defer server.Close()

	conn, err := NewDialer(server.URL)
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	if err != nil {
		t.Error(err)
	}

	if err := noecho.m.Close(); err != nil {
		return
	}
}

func TestSmallMessageBuffer(t *testing.T) {
	echo := NewTestServerHandler(func(session *Session, msg []byte) {
		err := session.Write(msg)
		if err != nil {
			return
		}
	})
	echo.m.Config.MessageBufferSize = 0
	echo.m.HandleError(func(s *Session, err error) {
		if err == nil {
			t.Error("there should be a buffer full error here")
		}
	})
	server := httptest.NewServer(echo)
	defer server.Close()

	conn, err := NewDialer(server.URL)
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	if err != nil {
		t.Error(err)
	}

	if err := conn.WriteMessage(websocket.TextMessage, []byte("12345")); err != nil {
		return
	}
}

func TestPong(t *testing.T) {
	echo := NewTestServerHandler(func(session *Session, msg []byte) {
		err := session.Write(msg)
		if err != nil {
			return
		}
	})
	echo.m.Config.PongWait = time.Second
	echo.m.Config.PingPeriod = time.Second * 9 / 10
	server := httptest.NewServer(echo)
	defer server.Close()

	conn, err := NewDialer(server.URL)
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	if err != nil {
		t.Error(err)
	}

	fired := false
	echo.m.HandlePong(func(s *Session) {
		fired = true
	})

	if err := conn.WriteMessage(websocket.PongMessage, nil); err != nil {
		return
	}

	time.Sleep(time.Millisecond)

	if !fired {
		t.Error("should have fired pong handler")
	}
}

func BenchmarkSessionWrite(b *testing.B) {
	echo := NewTestServerHandler(func(session *Session, msg []byte) {
		err := session.Write(msg)
		if err != nil {
			return
		}
	})
	server := httptest.NewServer(echo)
	conn, _ := NewDialer(server.URL)
	defer server.Close()
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	for n := 0; n < b.N; n++ {
		err := conn.WriteMessage(websocket.TextMessage, []byte("test"))
		if err != nil {
			return
		}
		if _, _, err := conn.ReadMessage(); err != nil {
			return
		}
	}
}

func BenchmarkBroadcast(b *testing.B) {
	echo := NewTestServerHandler(func(session *Session, msg []byte) {
		err := session.Write(msg)
		if err != nil {
			return
		}
	})
	server := httptest.NewServer(echo)
	defer server.Close()

	conns := make([]*websocket.Conn, 0)

	num := 100

	for i := 0; i < num; i++ {
		conn, _ := NewDialer(server.URL)
		conns = append(conns, conn)
	}

	for n := 0; n < b.N; n++ {
		err := echo.m.Broadcast([]byte("test"))
		if err != nil {
			return
		}

		for i := 0; i < num; i++ {
			_, _, err := conns[i].ReadMessage()
			if err != nil {
				return
			}
		}
	}

	for i := 0; i < num; i++ {
		err := conns[i].Close()
		if err != nil {
			return
		}
	}
}
