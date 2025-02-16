import NaturalRouter from 'natural-framework/router'
import * as HttpServer from 'natural-framework/server/uws'
// import * as HttpServer from 'natural-framework/server/node' // for Node.js native application

function createRoutes (router) {
  router
    .get('/', (_, response) => {
      response.end('')
    })
    .get('/user/:id', (request, response) => {
      response.end(request.params.id)
    })
    .post('/user', (request, response) => {
      response.end('')
    })
    .route({
      url: '/test/simple/:id',
      method: 'GET',
      type: 'json',
      handler: (request, response) => {
        // request.params, request.query, request.body, request.files
        response.send({ id: request.params.id }, 'json')
      }
    })

  /* router.route({
    url: '/meet/auth',
    method: 'GET',
    type: 'json',
    handler: (request, response) => {
      const params = Object.assign({}, request.params, request.query)
      response.send(params)
    }
  }) */

  // or
  /*
  router.on('GET', '/station/test/simple/:id', (request, response) => {
    // request.params, request.query, request.body, request.files
    response.send({ id: request.params.id }, 'json')
  })
  */
}

async function bootstrap () {
  const router = new NaturalRouter({
    server: HttpServer
    /* ssl: {
      key: path.join(__dirname, './security/cert.key'),
      cert: path.join(__dirname, './security/cert.pem')
    } */
  })
  try {
    createRoutes(router)
    const port = await router.listen(3000)
    console.log(`Listen http://localhost:${port}`)
  } catch (error) {
    console.log('Error:', error)
  }
}

bootstrap()
