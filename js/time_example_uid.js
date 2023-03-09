import { HawkFlow } from 'hawkflow';
import { setTimeout } from 'timers/promises';

// Authenticate with your API key
const hf = new HawkFlow({
    apiKey: 'YOUR_API_KEY',
    debug: true,
});

// Pass a unique string if you are sending through process and meta that
// are not unique. this could be for example if you are timing part of a distributed architecture
// where several instances of the same code is running on many machines or pat of a web app
// where many users will be hitting the same code at the same time
let uid = '123e4567-e89b-12d3-a456-426614174000';

// Start timing your code - pass through process (required) and meta (optional) parameters
console.log('Sending timing start data to hawkflow');
// Processing the response is optional
let rspStart = await hf.start({
    process: 'your_process',
    meta: 'your meta data',
    uid: uid,
});
console.log(rspStart);

// Do some work
await setTimeout(5000);

// End timing this piece of code - process (required) and meta (optional) parameters should match the start
console.log('Sending timing end data to hawkflow');
// Processing the response is optional
let rspEnd = await hf.end({
    process: 'your_process',
    meta: 'your meta data',
});
console.log(rspEnd);

