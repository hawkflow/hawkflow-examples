<?php

require 'vendor/autoload.php';

use HawkFlow\HawkFlow\HawkFlow;

// Authenticate with your API key
$hawkFlow = new HawkFlow("YOUR_API_KEY", 3, 1, true);

// Pass a unique string if you are sending through process and meta that
// are not unique. this could be for example if you are timing part of a distributed architecture
// where several instances of the same code is running on many machines or pat of a web app
// where many users will be hitting the same code at the same time
$uid = \uniqid();

// Start timing your code - pass through process (required) and meta (optional) parameters
echo "Sending timing start data to hawkflow\n";
// Processing the response is optional
$apiResponse = $hawkFlow->start("your_process_name", "your_meta_data", $uid);
echo $apiResponse . "\n";

// Do some work
echo "Sleeping for 5 seconds...\n";
sleep(5);

// End timing this piece of code - process (required) and meta (optional) parameters should match the start
echo "Sending timing end data to hawkflow\n";
// Processing the response is optional
$apiResponse = $hawkFlow->end("your_process_name", "your_meta_data", $uid);
echo $apiResponse . "\n";
