FROM nimlang/nim:onbuild

RUN nimble refresh

WORKDIR /superapp

COPY . .

RUN nimble build -y


# jester
EXPOSE 5000
# httpx
EXPOSE 8080

ENTRYPOINT ["./superapp"]