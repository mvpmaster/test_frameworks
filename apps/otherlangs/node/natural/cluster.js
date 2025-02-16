//server.js
const http = require("http");
const os = require("os");
//const cluster = require("cluster");
const cpuCores = 2 //os.cpus().length;
let test = 'test'.repeat(5000)
// if (cluster.isMaster) {
//   let instance = 0;
//   while (instance < cpuCores) {
//     cluster.fork();
//     ++instance;
//   }
// } else {
  console.log(`Child-process ${process.pid} started`);
  http.createServer((req, res) => {
    res.writeHead(200);
    res.end(test) //"server endpoint hit");
  }).listen(3000);
//}