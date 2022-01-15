import os 
input_path =  os.path.join(os.path.dirname(os.path.realpath(__file__)),'input.txt')

def min_number_of_ways(containers,total_amount):
    if total_amount == 0:
        return {"ways":1,"cost":0}
    if len(containers) == 0:
        return {"ways":0,"cost": float('inf') }

    container = containers[0]
    result_with_cont = {"ways":0,"cost":float('inf') }

    if total_amount - container >= 0:
        result_with_cont = min_number_of_ways(containers[1:],total_amount - container)
        result_with_cont["cost"] = result_with_cont["cost"] + 1 

    result_without_cont = min_number_of_ways(containers[1:],total_amount)

    if result_without_cont["cost"] ==  result_with_cont["cost"]:
        total_ways = result_without_cont["ways"] + result_with_cont["ways"]
        return {"ways": total_ways,"cost": result_with_cont["cost"]  }
    elif result_without_cont["cost"] < result_with_cont["cost"] :
        return result_without_cont
    else:
        return result_with_cont





def number_of_ways(containers,total_amount):
    if total_amount == 0:
        return 1
    if len(containers) == 0:
        return 0
    container = containers[0]
    ways_using_cont = 0
    if total_amount - container >= 0:
        ways_using_cont = number_of_ways(containers[1:],total_amount - container)

    return number_of_ways(containers[1:],total_amount) +  ways_using_cont

def main():
    input_liters = 150
    with open(input_path) as f:
        containers = [ int(s) for s in f.read().splitlines()]
        result1 = number_of_ways(containers,input_liters)
        result2 = min_number_of_ways(containers,input_liters)["ways"]

        print(f'the solution to part 1 is {result1}')
        print(f'the solution to part 2 is {result2}')




main()