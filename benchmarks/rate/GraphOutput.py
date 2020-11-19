#unfinished file
import matplotlib.pyplot as plt
import pandas as pd
import numpy as np
from statistics import mean 
import os
def main():
    #_______format cvs_________
    df = pd.read_csv(r'resFineGrained.txt', delim_whitespace=True, skiprows=4, header=None)
    df = df[:-2] #remove non relevant two last rows
    colsToRemove = [1,3,5,7]
    df.drop(df.columns[colsToRemove],axis=1,inplace=True)
    df.columns = ['ratio', 'ns/op','num_of_keys','num_of_total_bukets']
    num_of_keys = df['num_of_keys'][5].astype(int) #TODO: make this line more elegant
    SizeA = df['num_of_total_bukets'][5].astype(int) #TODO: make this line more elegant
    df['ratio'] =df['ratio'].str.lstrip("BenchmarkRatioRate/testAnchor_with_ratio_") #remove prefix
    df['ratio']=df['ratio'].astype(int)
    df['ns/op']=df['ns/op'].astype(int)
    #df['col_name'] = df['col_name'].str.replace('G', '1')
    #_______format graph________
    #gather all tests
    dfResults = df.groupby('ratio')['ns/op'].apply(list)
    #taking mean
    for i, row in dfResults.items():
          dfResults[i] = (mean(row)*(10**-9))/num_of_keys
    #find theortical results 
    ratios = dfResults.axes[0].tolist() #get all ratios
    baseline = dfResults.iloc[0]
    theoreticalResults = []
    for ratio in ratios:
        theoreticalResults.append(baseline*(1+np.log(ratio)))
    #plot theortical results
    plt.plot(ratios,theoreticalResults, label='Theoretical results', linestyle='--', marker='o', color='b')
    #show plot
    #dfResults['ns/op'].rename_axis
    plt.title("Anchor Hash with |A|=" + SizeA.astype(str))#+ "with diffrent Ratio's (Ratio |A|/|W|)")
    plt.ylabel("Seconds per key [sec]")
    dfResults.plot.line(label='Empirical results', linestyle='--', marker='o', color='r')
    plt.xlabel("Ratio |A|/|W|")
    plt.xscale("log") #log scale
    plt.legend()
    #plt.show()
    plt.savefig('results.png')



if __name__ == "__main__":
    main()


'''
ns = [2**x for x in range(11)];
data = {
  'goroutines': [6919, 13212, 25469, 50819, 88566, 162391, 299955, 574043, 1129372, 2251411, 4760560],
  'reflection': [10868, 22335, 54882, 148218, 543921, 1694021, 6102920, 22648976, 90204929, 383579039, 1676544681],
  'recursion': [2658, 14707, 44520, 114676, 261880, 560284, 1117642, 2242910, 4784719, 10044186, 20599475],
}
for (label, values) in data.items():
    plt.plot(ns, values, label=label)
plt.legend()
'''