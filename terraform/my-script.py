import sys
import json

def handler(event, data):
   data_as_json = json.loads(data)
   file = open("log.txt", "w+")
   print(event, file=file)
   file.flush()
   file.close()
   return {"id": "2", "name": "hello", "executor": "java"}
   
def read_data():
    data = ''
    for line in sys.stdin:
        data += line

    return data
   
if __name__ == '__main__':
    context = read_data()
    print(json.dumps(handler(sys.argv[1], context)))
