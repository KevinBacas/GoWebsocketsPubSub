(function() {
  'use strict';

  var int = 0;
  var startDummyWebSocket = function () {
    var sendMessage = function(ws) {
      ws.send(JSON.stringify({
        Message: (int).toString()
      }));
      int++;
    };

    var ws = new WebSocket("ws://localhost:8000/echo");
    ws.onopen = function() {
      sendMessage(this);
    };
    ws.onmessage = function(message) {
      console.log("Message received :", JSON.parse(message.data));
      setTimeout(function () {
        sendMessage(ws);
      }, 1000);
    };
    ws.onerror = function(error) {
      console.error("An error occured :", error.data);
    };
    ws.onclose = function() {
      console.error("Connection closed..");
    };

    console.log("Dummy WebSocket started");
  }

  for(var i = 0; i < 1; i++) {
    startDummyWebSocket();
  }

  console.log("Script started!");
})();
