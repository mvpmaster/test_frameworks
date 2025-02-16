https://phoenixnap.com/kb/install-ruby-ubuntu

rbenv -v
rbenv install -l
rbenv install [version number]
rbenv install 3.1.0
rbenv local 3.1.0
rbenv global 3.1.0

sudo apt-get install ruby-dev

----------------------

https://github.com/rbenv/ruby-build

git clone https://github.com/rbenv/ruby-build.git "$(rbenv root)"/plugins/ruby-build
git -C "$(rbenv root)"/plugins/ruby-build pull

https://stackoverflow.com/questions/29145108/how-to-install-ruby-on-linux-in-a-specific-folder

----------------------

rbenv install 3.1.0
wget https://cache.ruby-lang.org/pub/ruby/3.1/ruby-3.1.0.tar.gz

tar -xf ruby-3.1.0.tar.gz
cd ruby-3.1.0
./configure 
sudo make && sudo make install

----------------------

get the latest "Stable Snapshot" from https://ftp.ruby-lang.org/pub/ruby/stable-snapshot.tar.gz and not the "Current stable"
sudo apt-get install libffi-dev
sudo apt-get install zlibc zlib1g zlib1g-dev
sudo apt-get install openssl
sudo apt-get install libssl0.9.8 [[[   first, find the latest version with : apt-cache search libssl | grep SSL  ]]]
sudo apt-get install ca-certificates
sudo apt-get install libssl-dev
sudo apt-get install libreadline-dev

----------------------

sudo gem install syro

ruby test.rb

ruckup

-----

sudo gem install hanami

hanami new bookshelf

bundle install

hanami server