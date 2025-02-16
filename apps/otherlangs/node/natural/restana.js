const restana = require('restana')

test = 'test'.repeat(5000)
test2 = '{"error":"'+test+'"}'

const service = restana()
service.get('/test', (req, res) => res.send(test2))//'Hello World!'))

service.start(3000);