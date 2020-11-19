#######################################################################################
#######################################################################################
##### Function for construction of a graph from the "result_adjlist.txt" data set #####
#######################################################################################
#######################################################################################

adj_matrix_from_adj_list = function(adj_list, ID_mana){
  final_graph = data.frame(lapply(adj_list, as.character), stringsAsFactors = FALSE)
  n_of_nodes = nrow(final_graph)
  final_graph_int = matrix(rep(0, nrow(final_graph) * ncol(final_graph)), ncol = ncol(final_graph))
  
  for (node in 1:n_of_nodes){
    tmp = which(final_graph == ID_mana$nodeID[node], arr.ind = TRUE)
    final_graph_int[tmp] = node
  }
  
  final_graph_int = final_graph_int[order(final_graph_int[, 1]), ]
  final_graph_int = final_graph_int[, -1]
  
  adj_matrix = matrix(rep(0, n_of_nodes^2), nrow = n_of_nodes)
  for(node in 1:n_of_nodes){
    neighbours = final_graph_int[node, ]
    for (n in neighbours){
      adj_matrix[node, n] = 1
      adj_matrix[n, node] = 1
    }
  }
  return(adj_matrix)
}


##################################################################################
##################################################################################
#####                                                                        #####
##### Function for calculating the minimum percentage of the sum of the mana #####
##### of all the neighbours among top M richest nodes.                       #####
#####                                                                        #####
##################################################################################
##################################################################################

min_max_mana = function(M, adj_matrix, mana){
  sum_of_neigh_mana = rep(0, M)
  for (i in 1:M){
    neigh = which(adj_matrix[i, ] == 1)
    sum_of_neigh_mana[i] = sum(mana[neigh])
  }
  return (min(sum_of_neigh_mana)/sum(mana))
}