'use strict';

var mqtt;

// const host = "test.mosquitto.org";
// const port = 8080;

const host = "127.0.0.1";
const port = 9001;
const cname = "noobygames-tinygo-1";

function onConnect() {
    console.log("Successfully connected to mqtt the broker");

    mqtt.subscribe("/noobygames/homeautomation/home/bedroom/light/status");

    handleOnConnect();
}

function onMessageArrived(message) {
    console.log("onMessageArrived:" + message.payloadString);
    handleMessage(message.payloadString);
}

function publish(topic, message, qol) {
    mqtt.send(topic, message, qol, false);
}

function onConnectionLost(err) {
    if (err.errorCode !== 0) {
        console.log("onConnectionLost:" + err.errorMessage);
    }

    ConnectToMQTT();
}

function ConnectToMQTT() {
    console.log("mqtt client: trying to connect to " + host + ":" + port);

    mqtt = new Paho.MQTT.Client(host, port, cname);
    var options = {
        timeout: 1,
        onSuccess: onConnect,
    };

    mqtt.onConnectionLost = onConnectionLost;
    mqtt.onMessageArrived = onMessageArrived;

    mqtt.connect(options);
}