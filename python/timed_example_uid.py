import time
import uuid

from hawkflowclient.hawkflow_api import *


# authenticate with your API key
hf = HawkflowAPI("YOUR_API_KEY")

# pass a unique string if you are sending through process and meta that
# are not unique. this could be for example if you are timing part of a distributed architecture
# where several instances of the same code is running on many machines or pat of a web app
# where many users will be hitting the same code at the same time
unique_id = str(uuid.uuid4())

# start timing your code - pass through process (required) and meta (optional) parameters
print("sending timing start data to hawkflow")
hf.start("your_process_name_uid", "your_meta_data", uid=unique_id)

print("sleeping for 5 seconds...")
time.sleep(5)

# end timing this piece of code - process (required) and meta (optional) parameters should match the start
print("sending timing end data to hawkflow")
hf.end("your_process_name_uid", "your_meta_data", uid=unique_id)
