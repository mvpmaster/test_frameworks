require "syro"

App = Syro.new do
    get do
      res.write "GET /"
    end
  
    post do
      res.write "POST /"
    end
  
    on "users" do
      on :id do
  
        # Captured values go to the inbox
        @user = User[inbox[:id]]
  
        get do
          res.write "GET /users/42"
        end
  
        put do
          res.write "PUT /users/42"
        end
  
        patch do
          res.write "PATCH /users/42"
        end
  
        delete do
          res.write "DELETE /users/42"
        end
      end
  
      get do
        res.write "GET /users"
      end
  
      post do
        res.write "POST /users"
      end
    end
  end