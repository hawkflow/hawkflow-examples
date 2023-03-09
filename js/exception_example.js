import { HawkFlow } from 'hawkflow';

// Authenticate with your API key
const hf = new HawkFlow({
    apiKey: 'YOUR_API_KEY',
});

try {
    throw new Error('test exception');

} catch(e) {
    // Catch and send errors through to hawkflow
    console.log('Sending exception data to hawkflow');

    // Processing the response is optional
    let rsp = await hf.exception({
        process: 'your_process',
        meta: 'your meta data',
        message: e,
    });
    console.log(rsp);
}
