import time

from hawkflowclient.hawkflow_decorators import *


@HawkflowTimed(api_key="YOUR_API_KEY")
def hawkflow_examples_decorator_no_metadata():
    # time this code, not passing through any metadata
    time.sleep(5)


@HawkflowTimed(api_key="YOUR_API_KEY")
def hawkflow_examples_decorator_with_metadata(hawkflow_meta="your meta data"):
    # time this code, use the hawkflow_meta parameter to pass through any metadata
    time.sleep(5)


hawkflow_examples_decorator_no_metadata()

hawkflow_examples_decorator_with_metadata(hawkflow_meta="this is my metadata")

