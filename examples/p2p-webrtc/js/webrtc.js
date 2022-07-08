/*
 *  Copyright (c) 2015 The WebRTC project authors. All Rights Reserved.
 *
 *  Use of this source code is governed by a BSD-style license
 *  that can be found in the LICENSE file in the root of the source
 *  tree.
 */

"use strict";

let peerConnection;

const servers = { iceServers: [{ urls: "stun:stun.l.google.com:19302" }] };

const localSdpBox = document.querySelector("textarea#localSdp");
const remoteSdpBox = document.querySelector("textarea#remoteSdp");
const startButton = document.querySelector("button#startButton");

startButton.onclick = initiateConnection;
remoteSdpBox.onchange = handleRemoteAnswer;
window.onload = initiateConnection;

function enableStartButton() {
  startButton.disabled = false;
}

function initiateConnection() {
  localSdpBox.placeholder = "";
  window.peerConnection = peerConnection = new RTCPeerConnection(servers);
  peerConnection.isInitiator = true;

  console.log("Created local peer connection object localConnection");

  // localConnection handlers
  // Listen for local ICE candidates on the local RTCPeerConnection
  peerConnection.onicecandidate = (e) => {
    onIceCandidate(peerConnection, e);
  };

  peerConnection.onsignalingstatechange = (e) => {
    console.log(`Signaling state changed to: ${peerConnection.signalingState}`);
  };

  // TODO: trickle ICE candidates
  peerConnection.onicegatheringstatechange = (e) => {
    console.log(
      `ICE gathering state changed:`,
      e.currentTarget.iceGatheringState
    );
  };

  peerConnection.dataChannel =
    peerConnection.createDataChannel("sendDataChannel");
  console.log("Created send data channel");

  setDataChannelListeners(peerConnection.dataChannel);

  peerConnection
    .createOffer()
    .then(gotDescription, onCreateSessionDescriptionError);
  startButton.disabled = true;
}

// Switch from initiator to callee
function switchToCallee() {
  console.log("Switching to callee");
  peerConnection.isInitiator = false;
  localSdpBox.disabled = false;
  localSdpBox.placeholder = "Accepting call...";
  startButton.disabled = true;
  window.peerConnection = peerConnection = new RTCPeerConnection(servers);
  peerConnection.ondatachannel = (e) => {
    console.log("Received data channel");
    peerConnection.dataChannel = e.channel;
    setDataChannelListeners(peerConnection.dataChannel);
  };
}

function onCreateSessionDescriptionError(error) {
  console.error("Failed to create session description: " + error.toString());
}

function closeDataChannels() {
  console.log("Closing data channels");
  peerConnection.dataChannel.close();
  console.log(
    "Closed data channel with label: " + peerConnection.dataChannel.label
  );
  peerConnection.close();
  peerConnection = null;
  console.log("Closed peer connections");
  startButton.disabled = false;
  sendButton.disabled = true;
  localSdpBox.value = "";
  remoteSdpBox.value = "";
  localSdpBox.disabled = true;
  enableStartButton();
}

function gotDescription(desc) {
  peerConnection.setLocalDescription(desc);
  console.log(`Offer from localConnection\n${desc.sdp}`);
}

function getOtherPc(pc) {
  return pc === peerConnection ? remoteConnection : peerConnection;
}

function getName(pc) {
  return pc === peerConnection ? "localPeerConnection" : "remotePeerConnection";
}

function onIceCandidate(pc, event) {
  displayLocalSDP(pc, event);
  console.log(
    `${getName(pc)} ICE candidate: ${
      event.candidate ? event.candidate.candidate : "(null)"
    }`,
    event.candidate
  );
}

function onAddIceCandidateSuccess() {
  console.log("AddIceCandidate success.");
}

function onAddIceCandidateError(error) {
  console.error(`Failed to add Ice Candidate: ${error.toString()}`);
}

function onDataChannelStateChange() {
  const readyState = peerConnection.dataChannel.readyState;
  console.log("Send channel state is: " + readyState);
  if (readyState === "open") {
    window.onConnectionOpen(peerConnection);
  }
}

function onDataChannelMessage(event) {
  console.log("Received message: " + event.data);
}

function setDataChannelListeners(dc) {
  dc.onopen = onDataChannelStateChange;
  dc.onclose = onDataChannelStateChange;
  dc.onmessage = onDataChannelMessage;
}

function displayLocalSDP(pc, event) {
  if (!event.candidate) return;
  const json = JSON.stringify(peerConnection.localDescription);
  console.log({ json });
  localSdpBox.value = json + "\n\n";
}

async function handleRemoteAnswer(e) {
  const offerJson = e.target.value;
  console.log(`Remote candidate: ${offerJson}`);
  const offer = JSON.parse(offerJson);

  if (offer.type === "offer") {
    switchToCallee();
    await peerConnection.setRemoteDescription(offer);
    const answer = await peerConnection.createAnswer();
    await peerConnection.setLocalDescription(answer);
    peerConnection.dataChannel = console.log({ answer });
    e.target.value = "";
    localSdpBox.value = JSON.stringify(answer) + "\n\n";
    return;
  }
  // Assume it's an answer
  await peerConnection.setRemoteDescription(offer);
}
