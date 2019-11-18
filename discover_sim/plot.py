import numpy as np
import matplotlib.pyplot as plt

xcol = 0
xlabel = "Time [seconds]"
xscale = 'linear'

# - - - - parameters unlikely to be changed - - -
ycol = 1 
zcol = 2  # last column in the csv file
folder = "data/"


def main():
    printKnownPeerAnalysis()
    printSentMsgAnalysis()
    printSentMsgPdfAnalysis()

def printKnownPeerAnalysis(): 
    filename = folder+'plot_knownPeerAnalysis'
    partPlot3("KnownPeerAnalysis", "knownPeerAnalysis", filename, "blue")

def printSentMsgAnalysis(): 
    filename = folder+'plot_sentMsgAnalysis'
    histoPlot2("sentMsgAnalysis", "sentMsgAnalysis", filename, "blue")
    plt.xlabel("# of sent messages")
    plt.ylabel("# of nodes")
    plt.savefig(filename+'.eps', format='eps')
    plt.clf()

def printSentMsgPdfAnalysis(): 
    filename = folder+'plot_sentMsgPdfAnalysis'
    partPlot2("sentMsgPdfAnalysis", "sentMsgPdfAnalysis", filename, "blue")
    plt.xlabel("# of sent messages")
    plt.ylabel("PDF")
    plt.savefig(filename+'.eps', format='eps')
    plt.clf()

def histoPlot2(type, file, filename, color):
    color = 'tab:blue'
    bandwidth = 1 
    x = loadDatafromRow(file, ycol)
    plt.hist(x, bins=range(int(np.amin(x)), int(np.amax(x)), bandwidth))
    np.savez(filename+"_"+type, x=x)

def partPlot2(type, file, filename, color):
    color = 'tab:blue'
    x = loadDatafromRow(file, xcol)
    y = loadDatafromRow(file, ycol)
    x, y = sort2vecs(x, y)
    plt.plot(x, y, color=color, linewidth=1)
    np.savez(filename+"_"+type, x=x, y=y)

def partPlot3(type, file, filename, color):
    x = loadDatafromRow(file, xcol)
    y = loadDatafromRow(file, ycol)
    z = loadDatafromRow(file, zcol)
    x, y, z = sort3vecs(x, y, z)
    
    fig, ax1 = plt.subplots()
    
    color = 'tab:blue'
    ax1.set_xlabel('Time [seconds]')
    ax1.set_ylabel('Nodes know all peers in network [%]', color=color)
    ax1.set_ylim([0, 100])
    ax1.plot(x, y, color=color)
    ax1.tick_params(axis='y', labelcolor=color)

    ax2 = ax1.twinx()  # instantiate a second axes that shares the same x-axis

    color = 'tab:red'
    ax2.set_ylabel('Avg # of known peers', color=color)  # we already handled the x-label with ax1
    ax2.plot(x, z, color=color)
    ax2.tick_params(axis='y', labelcolor=color)

    fig.tight_layout()  # otherwise the right y-label is slightly clipped

    np.savez(filename+"_"+type, x=x, y=y)
    fig.savefig(filename+'.eps', format='eps')
    fig.clf()


def sort2vecs(x, y):
    i = np.argsort(x)
    x = x[i]
    y = y[i]
    return x, y

def sort3vecs(x, y, z):
    i = np.argsort(x)
    x = x[i]
    y = y[i]
    z = z[i]
    return x, y, z

def loadDatafromRow(datatype, row):
    try:
        filestr = folder+'result_'+datatype+'.csv'
        f = open(filestr, "r")
        data = np.loadtxt(f, delimiter=",", skiprows=1, usecols=(row))
        return data
    except IOError:
        print(filestr)
        print("File not found.")
        return []


# needs to be at the very end of the file
if __name__ == '__main__':
    main()