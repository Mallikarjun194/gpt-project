# """
# At the command line, only need to run once to install the package via pip:
#
# $ pip install google-generativeai
# """
#
# import google.generativeai as genai1
#
# genai1.configure(api_key="YOUR_API_KEY")
#
# # Set up the model
# generation_config = {
#     "temperature": 0.9,
#     "top_p": 1,
#     "top_k": 1,
#     "max_output_tokens": 2048,
# }
#
# safety_settings = [
#     {
#         "category": "HARM_CATEGORY_HARASSMENT",
#         "threshold": "BLOCK_MEDIUM_AND_ABOVE"
#     },
#     {
#         "category": "HARM_CATEGORY_HATE_SPEECH",
#         "threshold": "BLOCK_MEDIUM_AND_ABOVE"
#     },
#     {
#         "category": "HARM_CATEGORY_SEXUALLY_EXPLICIT",
#         "threshold": "BLOCK_MEDIUM_AND_ABOVE"
#     },
#     {
#         "category": "HARM_CATEGORY_DANGEROUS_CONTENT",
#         "threshold": "BLOCK_MEDIUM_AND_ABOVE"
#     },
# ]
#
# model = genai1.GenerativeModel(model_name="gemini-pro",
#                               generation_config=generation_config,
#                               safety_settings=safety_settings)
#
# prompt_parts = [
#     "input: Beauty\nShow me some perfumes\nShow me some fragrance\nCould you suggest some beauty products?\n#Beauty #perfume #LipEssentials",
#     "output: Tease Collector's Edition Eau De Parfum\nTease Mini Fragrance Duo\nDeluxe Mini Fragrance Trio\nTease Eau de Parfum\nTravel Fine Fragrance Mist\nTease Crème Cloud Eau de Parfum",
#     "input: Show me some lip essentials for Valentine's Day\nLip bomb\n#beauty #LipEssentials",
#     "output: Flavored Lip Gloss\nColor Shine Lip Gloss\nShine Plumper Extreme Lip Plumper\nLip Oil\nLip Care Kit",
#     "input: Nightwear\nShow me some trending regular Sleepwear\nSleepwear\nCould you show me some trending regular Nightwear?\nPlease show me some trending regular Nightwear\nNightDress\n#sleepshirt #camisets #pajamas",
#     "output: Short Cozy Robe\nSatin Jacquard Long Pajama Set\nModal Short Pajama Set\nThermal Sleepshirt\nSatin Long Pajama Set\nShimmer Knit Long Pajama Set",
#     "input: #beauty",
#     "output: ",
# ]
#
# response = model.generate_content(prompt_parts)
# print(response.text)