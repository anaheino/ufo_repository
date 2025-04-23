import numpy as np
from sklearn.cluster import KMeans
import matplotlib.pyplot as plt

# Example data: A list of coordinates (x, y)
coordinates = np.array([
    [1, 2], [2, 3], [3, 4], [5, 8], [8, 8],
    [1, 0.6], [9, 11], [8, 2], [10, 2], [9, 3]
])

# Number of clusters
k = 3

# Initialize and fit the KMeans algorithm
kmeans = KMeans(n_clusters=k)
kmeans.fit(coordinates)

# Get the cluster centers
centers = kmeans.cluster_centers_

# Get the labels of each point
labels = kmeans.labels_

# Plot the clusters
plt.scatter(coordinates[:, 0], coordinates[:, 1], c=labels, cmap='rainbow')
plt.scatter(centers[:, 0], centers[:, 1], color='black', marker='x', s=100)
plt.title('K-Means Clustering')
plt.xlabel('X Coordinate')
plt.ylabel('Y Coordinate')
plt.show()
