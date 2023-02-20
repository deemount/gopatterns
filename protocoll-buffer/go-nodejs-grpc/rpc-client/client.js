const GRPCClient = require('node-grpc-client')
const path = require("path")
const PROTO_PATH = path.resolve(__dirname, 'chat.proto')
const client = new GRPCClient(PROTO_PATH, 'chat', 'ChatService', 'localhost:9000')

const message = {
    body: "Hello from Salvatore",
}

const details = {
    name: "Salvatore",
    age: 44
}

client.runService('SayHello', message, (err, res) => {
    console.log(res.body);
});

client.runService('GetDetails', details, (err, res) => {
    console.log(res.body);
});