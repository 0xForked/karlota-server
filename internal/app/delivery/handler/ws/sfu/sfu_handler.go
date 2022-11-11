package sfu

//type websocketMessage struct {
//	Event string `json:"event"`
//	Data  string `json:"data"`
//}

// Helper to make Gorilla Websockets thread-safe
//type threadSafeWriter struct {
//	*websocket.Conn
//	sync.Mutex
//}

//type peerConnectionState struct {
//	peerConnection *webrtc.PeerConnection
//	websocket      *threadSafeWriter
//}

//var (
//listLock        sync.RWMutex
//peerConnections []peerConnectionState
//trackLocals     map[string]*webrtc.TrackLocalStaticRTP
//)

// Add to list of tracks and fire renegotiation for all PeerConnections
//func addTrack(t *webrtc.TrackRemote) *webrtc.TrackLocalStaticRTP {
//	listLock.Lock()
//	defer func() {
//		listLock.Unlock()
//		signalPeerConnections()
//	}()
//
//	// Create a new TrackLocal with the same codec as our incoming
//	trackLocal, err := webrtc.NewTrackLocalStaticRTP(t.Codec().RTPCodecCapability, t.ID(), t.StreamID())
//	if err != nil {
//		panic(err)
//	}
//
//	trackLocals[t.ID()] = trackLocal
//	return trackLocal
//}

// Remove from list of tracks and fire renegotiation for all PeerConnections
//func removeTrack(t *webrtc.TrackLocalStaticRTP) {
//	listLock.Lock()
//	defer func() {
//		listLock.Unlock()
//		signalPeerConnections()
//	}()
//
//	delete(trackLocals, t.ID())
//}

// signalPeerConnections updates each PeerConnection so that it is getting all the expected media tracks
//func signalPeerConnections() {
//	listLock.Lock()
//	defer func() {
//		listLock.Unlock()
//		dispatchKeyFrame()
//	}()
//
//	attemptSync := func() (tryAgain bool) {
//		for i := range peerConnections {
//			if peerConnections[i].peerConnection.ConnectionState() == webrtc.PeerConnectionStateClosed {
//				peerConnections = append(peerConnections[:i], peerConnections[i+1:]...)
//				return true // We modified the slice, start from the beginning
//			}
//
//			// map of sender we already are sending, so we don't double send
//			existingSenders := map[string]bool{}
//
//			for _, sender := range peerConnections[i].peerConnection.GetSenders() {
//				if sender.Track() == nil {
//					continue
//				}
//
//				existingSenders[sender.Track().ID()] = true
//
//				// If we have a RTPSender that doesn't map to an existing track remove and signal
//				if _, ok := trackLocals[sender.Track().ID()]; !ok {
//					if err := peerConnections[i].peerConnection.RemoveTrack(sender); err != nil {
//						return true
//					}
//				}
//			}
//
//			// Don't receive videos we are sending, make sure we don't have loop-back
//			for _, receiver := range peerConnections[i].peerConnection.GetReceivers() {
//				if receiver.Track() == nil {
//					continue
//				}
//
//				existingSenders[receiver.Track().ID()] = true
//			}
//
//			// Add all track we aren't sending yet to the PeerConnection
//			for trackID := range trackLocals {
//				if _, ok := existingSenders[trackID]; !ok {
//					if _, err := peerConnections[i].peerConnection.AddTrack(trackLocals[trackID]); err != nil {
//						return true
//					}
//				}
//			}
//
//			offer, err := peerConnections[i].peerConnection.CreateOffer(nil)
//			if err != nil {
//				return true
//			}
//
//			if err = peerConnections[i].peerConnection.SetLocalDescription(offer); err != nil {
//				return true
//			}
//
//			offerString, err := json.Marshal(offer)
//			if err != nil {
//				return true
//			}
//
//			if err = peerConnections[i].websocket.WriteJSON(&websocketMessage{
//				Event: "offer",
//				Data:  string(offerString),
//			}); err != nil {
//				return true
//			}
//		}
//
//		return
//	}
//
//	for syncAttempt := 0; ; syncAttempt++ {
//		if syncAttempt == 25 {
//			// Release the lock and attempt a sync in 3 seconds. We might be blocking a RemoveTrack or AddTrack
//			go func() {
//				time.Sleep(time.Second * 3)
//				signalPeerConnections()
//			}()
//			return
//		}
//
//		if !attemptSync() {
//			break
//		}
//	}
//}

// dispatchKeyFrame sends a keyframe to all PeerConnections, used everytime a new user joins the call
//func dispatchKeyFrame() {
//	listLock.Lock()
//	defer listLock.Unlock()
//
//	for i := range peerConnections {
//		for _, receiver := range peerConnections[i].peerConnection.GetReceivers() {
//			if receiver.Track() == nil {
//				continue
//			}
//
//			_ = peerConnections[i].peerConnection.WriteRTCP([]rtcp.Packet{
//				&rtcp.PictureLossIndication{
//					MediaSSRC: uint32(receiver.Track().SSRC()),
//				},
//			})
//		}
//	}
//}
