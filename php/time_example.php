<?php

require 'vendor/autoload.php';

use HawkFlow\HawkFlow\HawkFlow;

// Authenticate with your API key
$hawkFlow = new HawkFlow("YOUR_API_KEY");

// Start timing your code - pass through process (required) and meta (optional) parameters
echo "Sending timing start data to hawkflow\n";
$hawkFlow->start("your_process_name", "your_meta_data");

// Do some work
echo "Sleeping for 5 seconds...\n";
sleep(5);

// End timing this piece of code - process (required) and meta (optional) parameters should match the start
echo "Sending timing end data to hawkflow\n";
$hawkFlow->end("your_process_name", "your_meta_data");
