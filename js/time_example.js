import { HawkFlow } from 'hawkflow';
import { setTimeout } from 'timers/promises';

// Authenticate with your API key
const hf = new HawkFlow({
    apiKey: 'YOUR_API_KEY',
});

// Start timing your code - pass through process (required) and meta (optional) parameters
console.log('Sending timing start data to hawkflow');
hf.start({
    process: 'your_process',
    meta: 'your meta data',
});

// Do some work
await setTimeout(5000);

// End timing this piece of code - process (required) and meta (optional) parameters should match the start
console.log('Sending timing end data to hawkflow');
hf.end({
    process: 'your_process',
    meta: 'your meta data',
});

