test_id = "0000"

with open('tests/in/'+test_id+'.txt') as f:
    data = f.readlines()[59:]


for line in data:
  list = line.strip().split()
  print(' '.join(list[:4]))
  print(list[4])