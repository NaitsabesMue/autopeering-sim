####################################################################################
####################################################################################
##### This is a code for the analysis of the data obtained from simulations    #####
##### with different values of parameteres R and rho. The goal of the analysis #####
##### is to determine what is the best choice of the parameters that we have,  #####
##### where "the best" means that for this particular choice of parameters,    #####
##### some function that we have defined is maximized or minimized. The goal   #####
##### function will include some desirable properties of the final graph like  #####
##### the diameter and the sum of mana of neighbours of high mana nodes.       #####
####################################################################################
####################################################################################

### Initial part - removing everything form the environment and loading the igraph package
rm(list = ls())
dir()
library(igraph)
library(dplyr)


### Setting working directory where all the files are and loading auxiliary functions
setwd("~/go/autopeering-sim/Rcode")
source("auxFunctions.R")

##################################################################################
##################################################################################
#####
##### We read the outputs of simulations from /data/ folder                               
#####
##################################################################################
##################################################################################

### We first set constants

### These are the values of N, k, s, R and rho
N = 100
k = 6
s = 1
R = 10
rho = 2

num_of_sim = 1

#### Reading the node ID and corresponding mana values
mana_filename = paste("../data/mana_list-", as.character(N), "-",
                      as.character(k), "-", 
                      as.character(s), "-",
                      as.character(R), "-", 
                      as.character(rho), "-",
                      as.character(num_of_sim), ".txt", sep = "")
ID_mana = read.table(mana_filename, header = FALSE, fill = TRUE)
colnames(ID_mana) <- c("nodeID", "mana") 
ID_mana = data.frame(lapply(ID_mana, as.character), stringsAsFactors = FALSE)
ID_mana <- mutate(ID_mana, mana = as.numeric(mana))
#### Reading the adjacency list         
adj_filename = paste("../data/result_adjlist-", as.character(N), "-",
                     as.character(k), "-", 
                     as.character(s), "-",
                     as.character(R), "-", 
                     as.character(rho), "-",
                     as.character(num_of_sim), ".txt", sep = "")
adj_list = read.table(adj_filename, header = FALSE, fill = TRUE)

adj_matrix = adj_matrix_from_adj_list(adj_list, ID_mana)
            
current_graph = graph_from_adjacency_matrix(adj_matrix, mode = "undirected")
diameter(current_graph, unconnected = TRUE)

M <- 10 
min_max_mana(M, adj_matrix, ID_mana$mana)
          
eclipse_percentage <- numeric(N)
for (M in 1:N)
{
  eclipse_percentage[M] <- min_max_mana(M, adj_matrix, ID_mana$mana)
}
min(eclipse_percentage)

plot(eclipse_percentage, type="l", xlab="M")

