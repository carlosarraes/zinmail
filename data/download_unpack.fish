#!/usr/bin/env fish
set url "http://www.cs.cmu.edu/~enron/enron_mail_20150507.tar.gz"
set file_name "enron_mail_20150507.tar.gz"

echo "Downloading file..."
curl -O $url

echo "Unpacking file..."
tar -xzf $file_name

echo "Done!"
