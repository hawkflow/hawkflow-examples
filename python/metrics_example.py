from hawkflowclient.hawkflow_api import *


# authenticate with your API key
hf = HawkflowAPI("YOUR_API_KEY")

my_metrics = {
    "users_count": 5465,
    "model_accuracy": 0.85,
    "rows_entered_db": 4654465
}

print("sending metrics data to hawkflow")
api_response = hf.metrics("your_process_name", "your_meta_data", my_metrics)
print(api_response)
