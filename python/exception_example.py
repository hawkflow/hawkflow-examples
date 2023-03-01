import traceback

from hawkflowclient.hawkflow_api import *


# authenticate with your API key
hf = HawkflowAPI("YOUR_API_KEY")

try:
    1/0
except Exception as err:
    # catch and send exceptions through to hawkflow
    print("sending exception data to hawkflow")
    hf.exception("your_process_name", "your_meta_data", traceback.format_exc())
