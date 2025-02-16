const rayo = require('rayo');

test = 'test'.repeat(5000)
test2 = '{"error":"'+test+'"}'

rayo({ port: 3000 })
  .get('/test/:test', (req, res) => res.end(test)) //`Hello ${req.params.user}`))
  .start();
