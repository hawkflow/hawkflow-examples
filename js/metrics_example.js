import { HawkFlow } from 'hawkflow';

// Authenticate with your API key
const hf = new HawkFlow({
    apiKey: 'YOUR_API_KEY',
});

let items = {
    'users_count':     5465,
    'model_accuracy':  0.85,
    'rows_entered_db': 4654465,
};

console.log('Sending metrics data to hawkflow');
// Processing the response is optional
let rsp = await hf.metrics({
    process: 'your_process',
    meta: 'your meta data',
    items: items,
});
console.log(rsp);

