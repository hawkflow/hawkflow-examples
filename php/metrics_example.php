<?php

require 'vendor/autoload.php';

use HawkFlow\HawkFlow\HawkFlow;

// Authenticate with your API key
$hawkFlow = new HawkFlow("YOUR_API_KEY");

$items = [
    "users_count" => 5465,
    "model_accuracy" => 0.85,
    "rows_entered_db" => 4654465,
];

echo "Sending metrics data to hawkflow\n";
// Processing the response is optional
$apiResponse = $hawkFlow->metrics($items, "your_process_name", "your_meta_data");
echo $apiResponse . "\n";
