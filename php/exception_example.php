<?php

require 'vendor/autoload.php';

use HawkFlow\HawkFlow\HawkFlow;

// Authenticate with your API key
$hawkFlow = new HawkFlow("YOUR_API_KEY");

try {
    throw new \Exception("test exception");
} catch(\Exception $e) {
    // Catch and send errors through to hawkflow
    echo "Sending error data to hawkflow\n";

    // Processing the response is optional
    $apiResponse = $hawkFlow->exception($e->getMessage(), "your_process_name", "your_meta_data");
    echo $apiResponse . "\n";
}
