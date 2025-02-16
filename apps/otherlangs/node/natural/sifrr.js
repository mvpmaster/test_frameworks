require('uWebSockets.js').SSLApp({

    /* There are more SSL options, cut for brevity */
    key_file_name: 'misc/key.pem',
    cert_file_name: 'misc/cert.pem',
    
  }).ws('/*', {
  
    /* There are many common helper features */
    idleTimeout: 32,
    maxBackpressure: 1024,
    maxPayloadLength: 512,
    compression: DEDICATED_COMPRESSOR_3KB,
  
    /* For brevity we skip the other events (upgrade, open, ping, pong, close) */
    message: (ws, message, isBinary) => {
      /* You can do app.publish('sensors/home/temperature', '22C') kind of pub/sub as well */
      
      /* Here we echo the message back, using compression if available */
      let ok = ws.send(message, isBinary, true);
    }
    
  }).get('/*', (res, req) => {
  
    /* It does Http as well */
    res.writeStatus('200 OK').writeHeader('IsExample', 'Yes').end('Hello there!');
    
  }).listen(9001, (listenSocket) => {
  
    if (listenSocket) {
      console.log('Listening to port 9001');
    }
    
  });

// const { Cluster, App } = require('@sifrr/server');
// // const app = new App();
// // do something on app
// // const app2 = new App();
// // // do something on app2
// // const cluster = new Cluster([
// //   {
// //     app: app,
// //     port: 12345
// //   },
// //   {
// //     app: app2,
// //     ports: [12346, 12347]
// //   }
// // ]);

// const app = new App();
// // app.get('/', res => {
// // //   writeHeaders(res, name, value); // single header
// // //   writeHeaders(res, {
// // //     name1: value1,
// // //     name2: value2
// // //   }); // multiple headers
// // });

// let test = 'test'.repeat(5000)

// app.get('/', (res, req) => {
//     //writeHeaders(res, 'Connection', 'keep-alive'); // single header
//     res.json().then(jsonBody =>"{'test':1}")
// });
// //'ABD'));

// const cluster = new Cluster([
//     {
//       app: app,
//       port: 3000
//     },
//   ]);

// // listen on all ports
// cluster.listen((uwsSocket, port) => {
//   // this = app for port
//   console.log(this, `is listening on port ${port}`);
// });

// // // close all ports
// // cluster.close();

// // // close specific port
// // cluster.close(port);
