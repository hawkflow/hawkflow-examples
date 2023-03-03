import time

from hawkflowclient.hawkflow_api import *


# authenticate with your API key
hf = HawkflowAPI("YOUR_API_KEY")

# start timing your code - pass through process (required) and meta (optional) parameters
print("sending timing start data to hawkflow")
api_response = hf.start("your_process_name", "your_meta_data")
print(api_response)

print("sleeping for 5 seconds...")
time.sleep(5)

# end timing this piece of code - process (required) and meta (optional) parameters should match the start
print("sending timing end data to hawkflow")
api_response = hf.end("your_process_name", "your_meta_data")
print(api_response)
