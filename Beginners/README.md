# 400 Go Programming Questions for Kubernetes Use Cases (Beginner Level)

This document lists 400 beginner-level programming questions for implementing Go programs that interact with Kubernetes, focusing on Kubernetes API libraries (`k8s.io/client-go`), resource management (creations, modifications, deletions), and basic API interactions, excluding basic Go concepts. Each question begins with "Write a program" and assumes a configured Kubernetes client. The questions cover a wide variety of Kubernetes resources and operations, suitable for developers new to Kubernetes programming but familiar with Go.

## Beginner Questions

1. Write a program to list all pods in the `default` namespace using `client-go`, printing each pod’s name and status phase to the console.
2. Write a program to create a pod named `nginx-pod` in the `default` namespace using `client-go`, running an `nginx` container with the latest image.
3. Write a program to delete a pod named `nginx-pod` in the `default` namespace using `client-go`, ensuring graceful termination.
4. Write a program to retrieve details of a pod named `nginx-pod` in the `default` namespace using `client-go`, printing its IP address and node name.
5. Write a program to list all deployments in the `default` namespace using `client-go`, printing each deployment’s name and replica count.
6. Write a program to create a deployment named `nginx-deployment` in the `default` namespace using `client-go`, with 3 replicas of an `nginx` container.
7. Write a program to update the replica count of a deployment named `nginx-deployment` to 5 in the `default` namespace using `client-go`.
8. Write a program to delete a deployment named `nginx-deployment` in the `default` namespace using `client-go`.
9. Write a program to list all services in the `default` namespace using `client-go`, printing each service’s name and cluster IP.
10. Write a program to create a ClusterIP service named `nginx-service` in the `default` namespace using `client-go`, exposing port 80 for pods labeled `app=nginx`.
11. Write a program to list all ConfigMaps in the `default` namespace using `client-go`, printing each ConfigMap’s name and data keys.
12. Write a program to create a ConfigMap named `app-config` in the `default` namespace using `client-go`, with a key-value pair `key=value`.
13. Write a program to update a ConfigMap named `app-config` in the `default` namespace using `client-go`, changing the value of `key` to `new-value`.
14. Write a program to delete a ConfigMap named `app-config` in the `default` namespace using `client-go`.
15. Write a program to list all Secrets in the `default` namespace using `client-go`, printing each Secret’s name and type.
16. Write a program to create a Secret named `app-secret` in the `default` namespace using `client-go`, with a base64-encoded `password` field.
17. Write a program to list all nodes in the cluster using `client-go`, printing each node’s name and operating system.
18. Write a program to retrieve details of a node named `node-1` using `client-go`, printing its allocatable CPU and storage capacity.
19. Write a program to list all namespaces in the cluster using `client-go`, printing each namespace’s name and status.
20. Write a program to create a namespace named `test-ns` in the cluster using `client-go`.
21. Write a program to delete a namespace named `test-ns` from the cluster using `client-go`.
22. Write a program to list all PersistentVolumeClaims in the `default` namespace using `client-go`, printing each PVC’s name and requested storage size.
23. Write a program to create a PersistentVolumeClaim named `app-pvc` in the `default` namespace using `client-go`, requesting 2Gi storage with `ReadWriteOnce` access mode.
24. Write a program to list all PersistentVolumes in the cluster using `client-go`, printing each PV’s name and capacity.
25. Write a program to list all jobs in the `default` namespace using `client-go`, printing each job’s name and completion status.
26. Write a program to create a job named `pi-job` in the `default` namespace using `client-go`, running a `perl` container to calculate pi.
27. Write a program to list all CronJobs in the `default` namespace using `client-go`, printing each CronJob’s name and schedule.
28. Write a program to create a CronJob named `hello-cron` in the `default` namespace using `client-go`, running a `busybox` container every minute to print "hello".
29. Write a program to list all events in the `default` namespace using `client-go`, printing each event’s involved object and message.
30. Write a program to list all RoleBindings in the `default` namespace using `client-go`, printing each RoleBinding’s name and associated role.
31. Write a program to create a RoleBinding named `app-rolebinding` in the `default` namespace using `client-go`, binding the `default` service account to a role named `app-role`.
32. Write a program to list all Roles in the `default` namespace using `client-go`, printing each Role’s name and allowed verbs.
33. Write a program to create a Role named `pod-reader` in the `default` namespace using `client-go`, allowing `get` and `list` operations on pods.
34. Write a program to list all ServiceAccounts in the `default` namespace using `client-go`, printing each ServiceAccount’s name.
35. Write a program to create a ServiceAccount named `app-sa` in the `default` namespace using `client-go`.
36. Write a program to list all Ingresses in the `default` namespace using `client-go`, printing each Ingress’s name and host rules.
37. Write a program to create an Ingress named `app-ingress` in the `default` namespace using `client-go`, routing `app.example.com` to a service named `nginx-service` on port 80.
38. Write a program to list all StatefulSets in the `default` namespace using `client-go`, printing each StatefulSet’s name and replica count.
39. Write a program to create a StatefulSet named `mysql-sts` in the `default` namespace using `client-go`, running 3 replicas of a `mysql:5.7` container.
40. Write a program to list all DaemonSets in the `default` namespace using `client-go`, printing each DaemonSet’s name and node selector.
41. Write a program to list all ReplicaSets in the `default` namespace using `client-go`, printing each ReplicaSet’s name and desired replicas.
42. Write a program to create a ReplicaSet named `app-rs` in the `default` namespace using `client-go`, with 2 replicas of an `nginx` container.
43. Write a program to delete a ReplicaSet named `app-rs` in the `default` namespace using `client-go`.
44. Write a program to update a ReplicaSet named `app-rs` in the `default` namespace using `client-go`, changing its replicas to 4.
45. Write a program to list all HorizontalPodAutoscalers in the `default` namespace using `client-go`, printing each HPA’s name and target resource.
46. Write a program to create a HorizontalPodAutoscaler named `app-hpa` in the `default` namespace using `client-go`, scaling a deployment based on CPU usage.
47. Write a program to delete a HorizontalPodAutoscaler named `app-hpa` in the `default` namespace using `client-go`.
48. Write a program to list all NetworkPolicies in the `default` namespace using `client-go`, printing each policy’s name and pod selector.
49. Write a program to create a NetworkPolicy named `app-netpol` in the `default` namespace using `client-go`, allowing ingress traffic to pods labeled `app=nginx`.
50. Write a program to delete a NetworkPolicy named `app-netpol` in the `default` namespace using `client-go`.
51. Write a program to list all ResourceQuotas in the `default` namespace using `client-go`, printing each quota’s name and CPU limits.
52. Write a program to create a ResourceQuota named `app-quota` in the `default` namespace using `client-go`, limiting CPU to 2 cores.
53. Write a program to delete a ResourceQuota named `app-quota` in the `default` namespace using `client-go`.
54. Write a program to list all LimitRanges in the `default` namespace using `client-go`, printing each range’s name and default limits.
55. Write a program to create a LimitRange named `app-limit` in the `default` namespace using `client-go`, setting default CPU limits for containers.
56. Write a program to delete a LimitRange named `app-limit` in the `default` namespace using `client-go`.
57. Write a program to list all pods in the `kube-system` namespace using `client-go`, printing each pod’s name and creation timestamp.
58. Write a program to create a pod named `busybox-pod` in the `monitoring` namespace using `client-go`, running a `busybox` container with a sleep command.
59. Write a program to delete a pod named `busybox-pod` in the `monitoring` namespace using `client-go`.
60. Write a program to retrieve details of a pod named `busybox-pod` in the `monitoring` namespace using `client-go`, printing its container image.
61. Write a program to list all deployments in the `kube-system` namespace using `client-go`, printing each deployment’s name and image version.
62. Write a program to create a deployment named `app-deployment` in the `monitoring` namespace using `client-go`, with 2 replicas of a `redis` container.
63. Write a program to update a deployment named `app-deployment` in the `monitoring` namespace using `client-go`, changing its image to `redis:6.0`.
64. Write a program to delete a deployment named `app-deployment` in the `monitoring` namespace using `client-go`.
65. Write a program to list all services in the `kube-system` namespace using `client-go`, printing each service’s name and port details.
66. Write a program to create a NodePort service named `app-service` in the `monitoring` namespace using `client-go`, exposing port 3000 for pods labeled `app=redis`.
67. Write a program to delete a service named `app-service` in the `monitoring` namespace using `client-go`.
68. Write a program to list all ConfigMaps in the `kube-system` namespace using `client-go`, printing each ConfigMap’s name and number of keys.
69. Write a program to create a ConfigMap named `db-config` in the `monitoring` namespace using `client-go`, with a key-value pair `db=prod`.
70. Write a program to update a ConfigMap named `db-config` in the `monitoring` namespace using `client-go`, adding a new key `env=prod`.
71. Write a program to delete a ConfigMap named `db-config` in the `monitoring` namespace using `client-go`.
72. Write a program to list all Secrets in the `kube-system` namespace using `client-go`, printing each Secret’s name and data size.
73. Write a program to create a Secret named `db-secret` in the `monitoring` namespace using `client-go`, with a base64-encoded `username` field.
74. Write a program to delete a Secret named `db-secret` in the `monitoring` namespace using `client-go`.
75. Write a program to list all PersistentVolumeClaims in the `kube-system` namespace using `client-go`, printing each PVC’s name and phase.
76. Write a program to create a PersistentVolumeClaim named `db-pvc` in the `monitoring` namespace using `client-go`, requesting 1Gi storage with `ReadWriteMany` access mode.
77. Write a program to delete a PersistentVolumeClaim named `db-pvc` in the `monitoring` namespace using `client-go`.
78. Write a program to list all jobs in the `kube-system` namespace using `client-go`, printing each job’s name and parallelism setting.
79. Write a program to create a job named `cleanup-job` in the `monitoring` namespace using `client-go`, running a `busybox` container to remove temporary files.
80. Write a program to delete a job named `cleanup-job` in the `monitoring` namespace using `client-go`.
81. Write a program to list all CronJobs in the `kube-system` namespace using `client-go`, printing each CronJob’s name and last schedule time.
82. Write a program to create a CronJob named `backup-cron` in the `monitoring` namespace using `client-go`, running a backup script every hour.
83. Write a program to delete a CronJob named `backup-cron` in the `monitoring` namespace using `client-go`.
84. Write a program to list all RoleBindings in the `kube-system` namespace using `client-go`, printing each RoleBinding’s name and subjects.
85. Write a program to create a RoleBinding named `db-rolebinding` in the `monitoring` namespace using `client-go`, binding a service account to a role named `db-role`.
86. Write a program to delete a RoleBinding named `db-rolebinding` in the `monitoring` namespace using `client-go`.
87. Write a program to list all Roles in the `kube-system` namespace using `client-go`, printing each Role’s name and resource types.
88. Write a program to create a Role named `config-reader` in the `monitoring` namespace using `client-go`, allowing `get` and `list` operations on ConfigMaps.
89. Write a program to delete a Role named `config-reader` in the `monitoring` namespace using `client-go`.
90. Write a program to list all ServiceAccounts in the `kube-system` namespace using `client-go`, printing each ServiceAccount’s name and creation time.
91. Write a program to create a ServiceAccount named `db-sa` in the `monitoring` namespace using `client-go`.
92. Write a program to delete a ServiceAccount named `db-sa` in the `monitoring` namespace using `client-go`.
93. Write a program to list all Ingresses in the `kube-system` namespace using `client-go`, printing each Ingress’s name and backend service.
94. Write a program to create an Ingress named `db-ingress` in the `monitoring` namespace using `client-go`, routing `db.example.com` to a service named `db-service`.
95. Write a program to delete an Ingress named `db-ingress` in the `monitoring` namespace using `client-go`.
96. Write a program to list all StatefulSets in the `kube-system` namespace using `client-go`, printing each StatefulSet’s name and service name.
97. Write a program to create a StatefulSet named `redis-sts` in the `monitoring` namespace using `client-go`, running 2 replicas of a `redis:6.0` container.
98. Write a program to delete a StatefulSet named `redis-sts` in the `monitoring` namespace using `client-go`.
99. Write a program to list all DaemonSets in the `kube-system` namespace using `client-go`, printing each DaemonSet’s name and tolerations.
100. Write a program to create a DaemonSet named `log-agent` in the `monitoring` namespace using `client-go`, running a logging agent on each node.
101. Write a program to watch pods in the `default` namespace using `client-go`, printing creation events with pod names and namespaces.
102. Write a program to watch deployments in the `default` namespace using `client-go`, printing update events with deployment names and replica counts.
103. Write a program to watch services in the `default` namespace using `client-go`, printing deletion events with service names and types.
104. Write a program to watch ConfigMaps in the `default` namespace using `client-go`, printing creation events with ConfigMap names and data keys.
105. Write a program to watch Secrets in the `default` namespace using `client-go`, printing update events with Secret names and data sizes.
106. Write a program to watch PersistentVolumeClaims in the `default` namespace using `client-go`, printing deletion events with PVC names and storage requests.
107. Write a program to watch jobs in the `default` namespace using `client-go`, printing creation events with job names and completion times.
108. Write a program to watch CronJobs in the `default` namespace using `client-go`, printing update events with CronJob names and schedules.
109. Write a program to watch RoleBindings in the `default` namespace using `client-go`, printing deletion events with RoleBinding names and roles.
110. Write a program to watch Ingresses in the `default` namespace using `client-go`, printing creation events with Ingress names and hosts.
111. Write a program to watch pods in the `kube-system` namespace using `client-go`, printing update events with pod names and status phases.
112. Write a program to watch deployments in the `monitoring` namespace using `client-go`, printing creation events with deployment names and images.
113. Write a program to watch services in the `kube-system` namespace using `client-go`, printing deletion events with service names and ports.
114. Write a program to watch ConfigMaps in the `monitoring` namespace using `client-go`, printing update events with ConfigMap names and key counts.
115. Write a program to watch Secrets in the `kube-system` namespace using `client-go`, printing creation events with Secret names and types.
116. Write a program to create a pod with an `emptyDir` volume in the `default` namespace using `client-go`, running an `nginx` container.
117. Write a program to create a pod with a ConfigMap volume in the `default` namespace using `client-go`, mounting `app-config` into an `nginx` container.
118. Write a program to create a pod with a Secret volume in the `default` namespace using `client-go`, mounting `app-secret` into a `redis` container.
119. Write a program to create a pod with resource limits in the `default` namespace using `client-go`, setting CPU and memory limits for an `nginx` container.
120. Write a program to create a pod with node affinity in the `default` namespace using `client-go`, scheduling an `nginx` container on nodes labeled `zone=us-east`.
121. Write a program to create a service with `LoadBalancer` type in the `default` namespace using `client-go`, exposing port 80 for pods labeled `app=nginx`.
122. Write a program to create a service with `ExternalName` type in the `default` namespace using `client-go`, pointing to `external.example.com`.
123. Write a program to create a deployment with pod anti-affinity in the `default` namespace using `client-go`, ensuring `nginx` pods spread across nodes.
124. Write a program to create a deployment with tolerations in the `default` namespace using `client-go`, allowing `nginx` pods on tainted nodes.
125. Write a program to create a StatefulSet with a PersistentVolumeClaim in the `default` namespace using `client-go`, running a `mysql` container.
126. Write a program to patch a pod’s labels in the `default` namespace using `client-go`, adding `env=prod` to `nginx-pod`.
127. Write a program to patch a service’s annotations in the `default` namespace using `client-go`, adding `owner=team` to `nginx-service`.
128. Write a program to patch a deployment’s environment variables in the `default` namespace using `client-go`, adding `ENV=prod` to `nginx-deployment`.
129. Write a program to patch a ConfigMap’s data in the `default` namespace using `client-go`, updating `key` to `new-value` in `app-config`.
130. Write a program to patch a Secret’s data in the `default` namespace using `client-go`, updating `password` in `app-secret`.
131. Write a program to list all pods in the `default` namespace with label `app=nginx` using `client-go`, printing their names.
132. Write a program to list all services in the `default` namespace with annotation `managed=true` using `client-go`, printing their names.
133. Write a program to list all deployments in the `default` namespace with label `env=prod` using `client-go`, printing their replica counts.
134. Write a program to list all ConfigMaps in the `default` namespace with label `app=web` using `client-go`, printing their data keys.
135. Write a program to list all Secrets in the `default` namespace with annotation `secure=true` using `client-go`, printing their types.
136. Write a program to create a pod with a liveness probe in the `default` namespace using `client-go`, checking an `nginx` container’s health.
137. Write a program to create a pod with a readiness probe in the `default` namespace using `client-go`, ensuring a `redis` container is ready.
138. Write a program to create a deployment with a rolling update strategy in the `default` namespace using `client-go`, deploying `nginx` containers.
139. Write a program to create a deployment with a recreate strategy in the `default` namespace using `client-go`, deploying `redis` containers.
140. Write a program to create a service with session affinity in the `default` namespace using `client-go`, enabling sticky sessions for `nginx` pods.
141. Write a program to list all pods in the `monitoring` namespace with label `app=prometheus` using `client-go`, printing their status.
142. Write a program to create a ConfigMap with multiple keys in the `monitoring` namespace using `client-go`, including `db=prod` and `env=test`.
143. Write a program to create a Secret with multiple keys in the `monitoring` namespace using `client-go`, including `username` and `password`.
144. Write a program to create a pod with pod affinity in the `monitoring` namespace using `client-go`, scheduling `nginx` pods near `redis` pods.
145. Write a program to create a deployment with resource requests in the `monitoring` namespace using `client-go`, setting CPU and memory for `redis` containers.
146. Write a program to patch a pod’s annotations in the `kube-system` namespace using `client-go`, adding `managed=true` to a pod.
147. Write a program to patch a deployment’s labels in the `kube-system` namespace using `client-go`, adding `app=core` to a deployment.
148. Write a program to patch a service’s port in the `monitoring` namespace using `client-go`, changing `app-service` port to 8080.
149. Write a program to delete all pods with label `app=old` in the `default` namespace using `client-go`.
150. Write a program to delete all services with annotation `deprecated=true` in the `default` namespace using `client-go`.
151. Write a program to create a pod with a `hostPath` volume in the `default` namespace using `client-go`, mounting a local directory into an `nginx` container.
152. Write a program to create a pod with an init container in the `default` namespace using `client-go`, running a setup script before an `nginx` container.
153. Write a program to create a deployment with multiple containers in the `default` namespace using `client-go`, running `nginx` and `redis` in the same pod.
154. Write a program to create a service with multiple ports in the `default` namespace using `client-go`, exposing ports 80 and 443 for `nginx` pods.
155. Write a program to list all PersistentVolumes with label `type=ssd` using `client-go`, printing their names and capacities.
156. Write a program to create a PersistentVolumeClaim with storage class in the `default` namespace using `client-go`, requesting fast storage.
157. Write a program to create a job with backoff limit in the `default` namespace using `client-go`, retrying a task up to 3 times.
158. Write a program to create a CronJob with concurrency policy in the `default` namespace using `client-go`, forbidding concurrent runs.
159. Write a program to list all events for a specific pod in the `default` namespace using `client-go`, printing their reasons and messages.
160. Write a program to create a Role with multiple resources in the `default` namespace using `client-go`, allowing operations on pods and services.
161. Write a program to create an Ingress with TLS in the `default` namespace using `client-go`, securing `app.example.com` with a Secret.
162. Write a program to create a StatefulSet with pod management policy in the `default` namespace using `client-go`, using parallel pod creation.
163. Write a program to create a DaemonSet with update strategy in the `default` namespace using `client-go`, using rolling updates for agents.
164. Write a program to watch nodes in the cluster using `client-go`, printing update events with node names and conditions.
165. Write a program to watch namespaces in the cluster using `client-go`, printing creation events with namespace names.
166. Write a program to list all pods in the `default` namespace sorted by creation time using `client-go`, printing their names.
167. Write a program to create a pod with a specific service account in the `default` namespace using `client-go`, using `app-sa` for an `nginx` pod.
168. Write a program to create a deployment with a specific node selector in the `default` namespace using `client-go`, targeting nodes labeled `disk=ssd`.
169. Write a program to create a service with an external IP in the `default` namespace using `client-go`, assigning a static IP to `nginx-service`.
170. Write a program to patch a pod’s resource limits in the `default` namespace using `client-go`, updating CPU limits for `nginx-pod`.
171. Write a program to list all deployments in the `default` namespace with replicas greater than 3 using `client-go`, printing their names.
172. Write a program to create a ConfigMap from a file in the `default` namespace using `client-go`, loading data from `config.yaml`.
173. Write a program to create a Secret from literals in the `default` namespace using `client-go`, setting `username` and `password` directly.
174. Write a program to create a pod with a projected volume in the `default` namespace using `client-go`, combining ConfigMap and Secret data.
175. Write a program to create a job with a deadline in the `default` namespace using `client-go`, limiting execution time to 300 seconds.
176. Write a program to create a CronJob with a history limit in the `default` namespace using `client-go`, retaining 3 successful jobs.
177. Write a program to list all RoleBindings for a specific service account in the `default` namespace using `client-go`, printing their roles.
178. Write a program to create an Ingress with path-based routing in the `default` namespace using `client-go`, routing `/api` to a backend service.
179. Write a program to create a StatefulSet with a custom service in the `default` namespace using `client-go`, linking to a headless service.
180. Write a program to create a DaemonSet with a specific toleration in the `default` namespace using `client-go`, tolerating node taints.
181. Write a program to watch PersistentVolumes in the cluster using `client-go`, printing deletion events with PV names.
182. Write a program to watch ResourceQuotas in the `default` namespace using `client-go`, printing update events with quota limits.
183. Write a program to list all pods in the `default` namespace with specific annotations using `client-go`, printing their names.
184. Write a program to create a pod with a specific priority class in the `default` namespace using `client-go`, using `high-priority` for `nginx`.
185. Write a program to create a deployment with a custom update strategy in the `default` namespace using `client-go`, setting max surge to 2.
186. Write a program to create a service with a specific cluster IP in the `default` namespace using `client-go`, assigning a fixed IP.
187. Write a program to patch a ConfigMap with a new key in the `default` namespace using `client-go`, adding `version=1.0` to `app-config`.
188. Write a program to patch a Secret with a new key in the `default` namespace using `client-go`, adding `token` to `app-secret`.
189. Write a program to create a PersistentVolumeClaim with annotations in the `default` namespace using `client-go`, adding `owner=app`.
190. Write a program to create a job with a specific namespace in the `monitoring` namespace using `client-go`, running a cleanup task.
191. Write a program to create a CronJob with a specific time zone in the `default` namespace using `client-go`, scheduling in UTC.
192. Write a program to list all Ingresses with specific annotations in the `default` namespace using `client-go`, printing their hosts.
193. Write a program to create a StatefulSet with a custom volume claim template in the `default` namespace using `client-go`, using dynamic PVCs.
194. Write a program to create a DaemonSet with a specific node affinity in the `default` namespace using `client-go`, targeting specific nodes.
195. Write a program to watch LimitRanges in the `default` namespace using `client-go`, printing creation events with limit details.
196. Write a program to watch NetworkPolicies in the `default` namespace using `client-go`, printing update events with policy details.
197. Write a program to list all services in the `default` namespace with specific labels using `client-go`, printing their ports.
198. Write a program to create a pod with a specific security context in the `default` namespace using `client-go`, running `nginx` as non-root.
199. Write a program to create a deployment with a specific pod template annotation in the `default` namespace using `client-go`, adding `version=1`.
200. Write a program to create a service with a specific external traffic policy in the `default` namespace using `client-go`, setting to `Local`.
201. Write a program to list all pods in the `app` namespace using `client-go`, printing each pod’s name and status phase to the console.
202. Write a program to create a pod named `redis-pod` in the `app` namespace using `client-go`, running a `redis` container with the latest image.
203. Write a program to delete a pod named `redis-pod` in the `app` namespace using `client-go`, ensuring graceful termination.
204. Write a program to retrieve details of a pod named `redis-pod` in the `app` namespace using `client-go`, printing its IP address and node name.
205. Write a program to list all deployments in the `app` namespace using `client-go`, printing each deployment’s name and replica count.
206. Write a program to create a deployment named `redis-deployment` in the `app` namespace using `client-go`, with 3 replicas of a `redis` container.
207. Write a program to update the replica count of a deployment named `redis-deployment` to 5 in the `app` namespace using `client-go`.
208. Write a program to delete a deployment named `redis-deployment` in the `app` namespace using `client-go`.
209. Write a program to list all services in the `app` namespace using `client-go`, printing each service’s name and cluster IP.
210. Write a program to create a ClusterIP service named `redis-service` in the `app` namespace using `client-go`, exposing port 6379 for pods labeled `app=redis`.
211. Write a program to list all ConfigMaps in the `app` namespace using `client-go`, printing each ConfigMap’s name and data keys.
212. Write a program to create a ConfigMap named `redis-config` in the `app` namespace using `client-go`, with a key-value pair `key=value`.
213. Write a program to update a ConfigMap named `redis-config` in the `app` namespace using `client-go`, changing the value of `key` to `new-value`.
214. Write a program to delete a ConfigMap named `redis-config` in the `app` namespace using `client-go`.
215. Write a program to list all Secrets in the `app` namespace using `client-go`, printing each Secret’s name and type.
216. Write a program to create a Secret named `redis-secret` in the `app` namespace using `client-go`, with a base64-encoded `password` field.
217. Write a program to list all PersistentVolumeClaims in the `app` namespace using `client-go`, printing each PVC’s name and requested storage size.
218. Write a program to create a PersistentVolumeClaim named `redis-pvc` in the `app` namespace using `client-go`, requesting 2Gi storage with `ReadWriteOnce` access mode.
219. Write a program to list all jobs in the `app` namespace using `client-go`, printing each job’s name and completion status.
220. Write a program to create a job named `redis-job` in the `app` namespace using `client-go`, running a `redis` container to perform a task.
221. Write a program to list all CronJobs in the `app` namespace using `client-go`, printing each CronJob’s name and schedule.
222. Write a program to create a CronJob named `redis-cron` in the `app` namespace using `client-go`, running a `redis` container every minute.
223. Write a program to list all RoleBindings in the `app` namespace using `client-go`, printing each RoleBinding’s name and associated role.
224. Write a program to create a RoleBinding named `redis-rolebinding` in the `app` namespace using `client-go`, binding the `default` service account to a role named `redis-role`.
225. Write a program to list all Roles in the `app` namespace using `client-go`, printing each Role’s name and allowed verbs.
226. Write a program to create a Role named `redis-reader` in the `app` namespace using `client-go`, allowing `get` and `list` operations on pods.
227. Write a program to list all ServiceAccounts in the `app` namespace using `client-go`, printing each ServiceAccount’s name.
228. Write a program to create a ServiceAccount named `redis-sa` in the `app` namespace using `client-go`.
229. Write a program to list all Ingresses in the `app` namespace using `client-go`, printing each Ingress’s name and host rules.
230. Write a program to create an Ingress named `redis-ingress` in the `app` namespace using `client-go`, routing `redis.example.com` to a service named `redis-service` on port 6379.
231. Write a program to list all StatefulSets in the `app` namespace using `client-go`, printing each StatefulSet’s name and replica count.
232. Write a program to create a StatefulSet named `redis-sts` in the `app` namespace using `client-go`, running 3 replicas of a `redis` container.
233. Write a program to list all DaemonSets in the `app` namespace using `client-go`, printing each DaemonSet’s name and node selector.
234. Write a program to create a DaemonSet named `redis-agent` in the `app` namespace using `client-go`, running a `redis` agent on each node.
235. Write a program to watch pods in the `app` namespace using `client-go`, printing creation events with pod names and namespaces.
236. Write a program to watch deployments in the `app` namespace using `client-go`, printing update events with deployment names and replica counts.
237. Write a program to watch services in the `app` namespace using `client-go`, printing deletion events with service names and types.
238. Write a program to watch ConfigMaps in the `app` namespace using `client-go`, printing creation events with ConfigMap names and data keys.
239. Write a program to watch Secrets in the `app` namespace using `client-go`, printing update events with Secret names and data sizes.
240. Write a program to create a pod with an `emptyDir` volume in the `app` namespace using `client-go`, running a `redis` container.
241. Write a program to create a pod with a ConfigMap volume in the `app` namespace using `client-go`, mounting `redis-config` into a `redis` container.
242. Write a program to create a pod with a Secret volume in the `app` namespace using `client-go`, mounting `redis-secret` into a `redis` container.
243. Write a program to create a pod with resource limits in the `app` namespace using `client-go`, setting CPU and memory limits for a `redis` container.
244. Write a program to create a pod with node affinity in the `app` namespace using `client-go`, scheduling a `redis` container on nodes labeled `zone=us-west`.
245. Write a program to create a service with `LoadBalancer` type in the `app` namespace using `client-go`, exposing port 6379 for pods labeled `app=redis`.
246. Write a program to create a deployment with pod anti-affinity in the `app` namespace using `client-go`, ensuring `redis` pods spread across nodes.
247. Write a program to create a deployment with tolerations in the `app` namespace using `client-go`, allowing `redis` pods on tainted nodes.
248. Write a program to create a StatefulSet with a PersistentVolumeClaim in the `app` namespace using `client-go`, running a `redis` container.
249. Write a program to patch a pod’s labels in the `app` namespace using `client-go`, adding `env=prod` to `redis-pod`.
250. Write a program to patch a service’s annotations in the `app` namespace using `client-go`, adding `owner=team` to `redis-service`.
251. Write a program to patch a deployment’s environment variables in the `app` namespace using `client-go`, adding `ENV=prod` to `redis-deployment`.
252. Write a program to patch a ConfigMap’s data in the `app` namespace using `client-go`, updating `key` to `new-value` in `redis-config`.
253. Write a program to patch a Secret’s data in the `app` namespace using `client-go`, updating `password` in `redis-secret`.
254. Write a program to list all pods in the `app` namespace with label `app=redis` using `client-go`, printing their names.
255. Write a program to list all services in the `app` namespace with annotation `managed=true` using `client-go`, printing their names.
256. Write a program to list all deployments in the `app` namespace with label `env=prod` using `client-go`, printing their replica counts.
257. Write a program to list all ConfigMaps in the `app` namespace with label `app=redis` using `client-go`, printing their data keys.
258. Write a program to list all Secrets in the `app` namespace with annotation `secure=true` using `client-go`, printing their types.
259. Write a program to create a pod with a liveness probe in the `app` namespace using `client-go`, checking a `redis` container’s health.
260. Write a program to create a pod with a readiness probe in the `app` namespace using `client-go`, ensuring a `redis` container is ready.
261. Write a program to create a deployment with a rolling update strategy in the `app` namespace using `client-go`, deploying `redis` containers.
262. Write a program to create a deployment with a recreate strategy in the `app` namespace using `client-go`, deploying `redis` containers.
263. Write a program to create a service with session affinity in the `app` namespace using `client-go`, enabling sticky sessions for `redis` pods.
264. Write a program to list all pods in the `app` namespace with label `app=redis` using `client-go`, printing their status.
265. Write a program to create a ConfigMap with multiple keys in the `app` namespace using `client-go`, including `db=prod` and `env=test`.
266. Write a program to create a Secret with multiple keys in the `app` namespace using `client-go`, including `username` and `password`.
267. Write a program to create a pod with pod affinity in the `app` namespace using `client-go`, scheduling `redis` pods near `nginx` pods.
268. Write a program to create a deployment with resource requests in the `app` namespace using `client-go`, setting CPU and memory for `redis` containers.
269. Write a program to patch a pod’s annotations in the `app` namespace using `client-go`, adding `managed=true` to a pod.
270. Write a program to patch a deployment’s labels in the `app` namespace using `client-go`, adding `app=core` to a deployment.
271. Write a program to patch a service’s port in the `app` namespace using `client-go`, changing `redis-service` port to 6380.
272. Write a program to delete all pods with label `app=old` in the `app` namespace using `client-go`.
273. Write a program to delete all services with annotation `deprecated=true` in the `app` namespace using `client-go`.
274. Write a program to create a pod with a `hostPath` volume in the `app` namespace using `client-go`, mounting a local directory into a `redis` container.
275. Write a program to create a pod with an init container in the `app` namespace using `client-go`, running a setup script before a `redis` container.
276. Write a program to create a deployment with multiple containers in the `app` namespace using `client-go`, running `redis` and `nginx` in the same pod.
277. Write a program to create a service with multiple ports in the `app` namespace using `client-go`, exposing ports 6379 and 8080 for `redis` pods.
278. Write a program to create a job with backoff limit in the `app` namespace using `client-go`, retrying a task up to 3 times.
279. Write a program to create a CronJob with concurrency policy in the `app` namespace using `client-go`, forbidding concurrent runs.
280. Write a program to list all events for a specific pod in the `app` namespace using `client-go`, printing their reasons and messages.
281. Write a program to create a Role with multiple resources in the `app` namespace using `client-go`, allowing operations on pods and services.
282. Write a program to create an Ingress with TLS in the `app` namespace using `client-go`, securing `redis.example.com` with a Secret.
283. Write a program to create a StatefulSet with pod management policy in the `app` namespace using `client-go`, using parallel pod creation.
284. Write a program to create a DaemonSet with update strategy in the `app` namespace using `client-go`, using rolling updates for agents.
285. Write a program to watch pods in the `app` namespace using `client-go`, printing deletion events with pod names and namespaces.
286. Write a program to watch deployments in the `app` namespace using `client-go`, printing creation events with deployment names and images.
287. Write a program to watch services in the `app` namespace using `client-go`, printing update events with service names and ports.
288. Write a program to watch ConfigMaps in the `app` namespace using `client-go`, printing deletion events with ConfigMap names and key counts.
289. Write a program to watch Secrets in the `app` namespace using `client-go`, printing creation events with Secret names and types.
290. Write a program to create a pod with a specific service account in the `app` namespace using `client-go`, using `redis-sa` for a `redis` pod.
291. Write a program to create a deployment with a specific node selector in the `app` namespace using `client-go`, targeting nodes labeled `disk=ssd`.
292. Write a program to create a service with an external IP in the `app` namespace using `client-go`, assigning a static IP to `redis-service`.
293. Write a program to patch a pod’s resource limits in the `app` namespace using `client-go`, updating CPU limits for `redis-pod`.
294. Write a program to list all deployments in the `app` namespace with replicas greater than 3 using `client-go`, printing their names.
295. Write a program to create a ConfigMap from a file in the `app` namespace using `client-go`, loading data from `config.yaml`.
296. Write a program to create a Secret from literals in the `app` namespace using `client-go`, setting `username` and `password` directly.
297. Write a program to create a pod with a projected volume in the `app` namespace using `client-go`, combining ConfigMap and Secret data.
298. Write a program to create a job with a deadline in the `app` namespace using `client-go`, limiting execution time to 300 seconds.
299. Write a program to create a CronJob with a history limit in the `app` namespace using `client-go`, retaining 3 successful jobs.
300. Write a program to list all RoleBindings for a specific service account in the `app` namespace using `client-go`, printing their roles.
301. Write a program to create an Ingress with path-based routing in the `app` namespace using `client-go`, routing `/api` to a backend service.
302. Write a program to create a StatefulSet with a custom service in the `app` namespace using `client-go`, linking to a headless service.
303. Write a program to create a DaemonSet with a specific toleration in the `app` namespace using `client-go`, tolerating node taints.
304. Write a program to watch PersistentVolumes in the cluster using `client-go`, printing deletion events with PV names.
305. Write a program to watch ResourceQuotas in the `app` namespace using `client-go`, printing update events with quota limits.
306. Write a program to list all pods in the `app` namespace with specific annotations using `client-go`, printing their names.
307. Write a program to create a pod with a specific priority class in the `app` namespace using `client-go`, using `high-priority` for `redis`.
308. Write a program to create a deployment with a custom update strategy in the `app` namespace using `client-go`, setting max surge to 2.
309. Write a program to create a service with a specific cluster IP in the `app` namespace using `client-go`, assigning a fixed IP.
310. Write a program to patch a ConfigMap with a new key in the `app` namespace using `client-go`, adding `version=1.0` to `redis-config`.
311. Write a program to patch a Secret with a new key in the `app` namespace using `client-go`, adding `token` to `redis-secret`.
312. Write a program to create a PersistentVolumeClaim with annotations in the `app` namespace using `client-go`, adding `owner=app`.
313. Write a program to create a job with a specific namespace in the `app` namespace using `client-go`, running a cleanup task.
314. Write a program to create a CronJob with a specific time zone in the `app` namespace using `client-go`, scheduling in UTC.
315. Write a program to list all Ingresses with specific annotations in the `app` namespace using `client-go`, printing their hosts.
316. Write a program to create a StatefulSet with a custom volume claim template in the `app` namespace using `client-go`, using dynamic PVCs.
317. Write a program to create a DaemonSet with a specific node affinity in the `app` namespace using `client-go`, targeting specific nodes.
318. Write a program to watch LimitRanges in the `app` namespace using `client-go`, printing creation events with limit details.
319. Write a program to watch NetworkPolicies in the `app` namespace using `client-go`, printing update events with policy details.
320. Write a program to list all services in the `app` namespace with specific labels using `client-go`, printing their ports.
321. Write a program to create a pod with a specific security context in the `app` namespace using `client-go`, running `redis` as non-root.
322. Write a program to create a deployment with a specific pod template annotation in the `app` namespace using `client-go`, adding `version=1`.
323. Write a program to create a service with a specific external traffic policy in the `app` namespace using `client-go`, setting to `Local`.
324. Write a program to list all pods in the `test` namespace using `client-go`, printing each pod’s name and status phase to the console.
325. Write a program to create a pod named `app-pod` in the `test` namespace using `client-go`, running an `nginx` container with the latest image.
326. Write a program to delete a pod named `app-pod` in the `test` namespace using `client-go`, ensuring graceful termination.
327. Write a program to retrieve details of a pod named `app-pod` in the `test` namespace using `client-go`, printing its IP address and node name.
328. Write a program to list all deployments in the `test` namespace using `client-go`, printing each deployment’s name and replica count.
329. Write a program to create a deployment named `app-deployment` in the `test` namespace using `client-go`, with 3 replicas of an `nginx` container.
330. Write a program to update the replica count of a deployment named `app-deployment` to 5 in the `test` namespace using `client-go`.
331. Write a program to delete a deployment named `app-deployment` in the `test` namespace using `client-go`.
332. Write a program to list all services in the `test` namespace using `client-go`, printing each service’s name and cluster IP.
333. Write a program to create a ClusterIP service named `app-service` in the `test` namespace using `client-go`, exposing port 80 for pods labeled `app=nginx`.
334. Write a program to list all ConfigMaps in the `test` namespace using `client-go`, printing each ConfigMap’s name and data keys.
335. Write a program to create a ConfigMap named `app-config` in the `test` namespace using `client-go`, with a key-value pair `key=value`.
336. Write a program to update a ConfigMap named `app-config` in the `test` namespace using `client-go`, changing the value of `key` to `new-value`.
337. Write a program to delete a ConfigMap named `app-config` in the `test` namespace using `client-go`.
338. Write a program to list all Secrets in the `test` namespace using `client-go`, printing each Secret’s name and type.
339. Write a program to create a Secret named `app-secret` in the `test` namespace using `client-go`, with a base64-encoded `password` field.
340. Write a program to list all PersistentVolumeClaims in the `test` namespace using `client-go`, printing each PVC’s name and requested storage size.
341. Write a program to create a PersistentVolumeClaim named `app-pvc` in the `test` namespace using `client-go`, requesting 2Gi storage with `ReadWriteOnce` access mode.
342. Write a program to list all jobs in the `test` namespace using `client-go`, printing each job’s name and completion status.
343. Write a program to create a job named `app-job` in the `test` namespace using `client-go`, running an `nginx` container to perform a task.
344. Write a program to list all CronJobs in the `test` namespace using `client-go`, printing each CronJob’s name and schedule.
345. Write a program to create a CronJob named `app-cron` in the `test` namespace using `client-go`, running an `nginx` container every minute.
346. Write a program to list all RoleBindings in the `test` namespace using `client-go`, printing each RoleBinding’s name and associated role.
347. Write a program to create a RoleBinding named `app-rolebinding` in the `test` namespace using `client-go`, binding the `default` service account to a role named `app-role`.
348. Write a program to list all Roles in the `test` namespace using `client-go`, printing each Role’s name and allowed verbs.
349. Write a program to create a Role named `app-reader` in the `test` namespace using `client-go`, allowing `get` and `list` operations on pods.
350. Write a program to list all ServiceAccounts in the `test` namespace using `client-go`, printing each ServiceAccount’s name.
351. Write a program to create a ServiceAccount named `app-sa` in the `test` namespace using `client-go`.
352. Write a program to list all Ingresses in the `test` namespace using `client-go`, printing each Ingress’s name and host rules.
353. Write a program to create an Ingress named `app-ingress` in the `test` namespace using `client-go`, routing `app.example.com` to a service named `app-service` on port 80.
354. Write a program to list all StatefulSets in the `test` namespace using `client-go`, printing each StatefulSet’s name and replica count.
355. Write a program to create a StatefulSet named `app-sts` in the `test` namespace using `client-go`, running 3 replicas of an `nginx` container.
356. Write a program to list all DaemonSets in the `test` namespace using `client-go`, printing each DaemonSet’s name and node selector.
357. Write a program to create a DaemonSet named `app-agent` in the `test` namespace using `client-go`, running an agent on each node.
358. Write a program to watch pods in the `test` namespace using `client-go`, printing creation events with pod names and namespaces.
359. Write a program to watch deployments in the `test` namespace using `client-go`, printing update events with deployment names and replica counts.
360. Write a program to watch services in the `test` namespace using `client-go`, printing deletion events with service names and types.
361. Write a program to watch ConfigMaps in the `test` namespace using `client-go`, printing creation events with ConfigMap names and data keys.
362. Write a program to watch Secrets in the `test` namespace using `client-go`, printing update events with Secret names and data sizes.
363. Write a program to create a pod with an `emptyDir` volume in the `test` namespace using `client-go`, running an `nginx` container.
364. Write a program to create a pod with a ConfigMap volume in the `test` namespace using `client-go`, mounting `app-config` into an `nginx` container.
365. Write a program to create a pod with a Secret volume in the `test` namespace using `client-go`, mounting `app-secret` into an `nginx` container.
366. Write a program to create a pod with resource limits in the `test` namespace using `client-go`, setting CPU and memory limits for an `nginx` container.
367. Write a program to create a pod with node affinity in the `test` namespace using `client-go`, scheduling an `nginx` container on nodes labeled `zone=us-east`.
368. Write a program to create a service with `LoadBalancer` type in the `test` namespace using `client-go`, exposing port 80 for pods labeled `app=nginx`.
369. Write a program to create a deployment with pod anti-affinity in the `test` namespace using `client-go`, ensuring `nginx` pods spread across nodes.
370. Write a program to create a deployment with tolerations in the `test` namespace using `client-go`, allowing `nginx` pods on tainted nodes.
371. Write a program to create a StatefulSet with a PersistentVolumeClaim in the `test` namespace using `client-go`, running an `nginx` container.
372. Write a program to patch a pod’s labels in the `test` namespace using `client-go`, adding `env=prod` to `app-pod`.
373. Write a program to patch a service’s annotations in the `test` namespace using `client-go`, adding `owner=team` to `app-service`.
374. Write a program to patch a deployment’s environment variables in the `test` namespace using `client-go`, adding `ENV=prod` to `app-deployment`.
375. Write a program to patch a ConfigMap’s data in the `test` namespace using `client-go`, updating `key` to `new-value` in `app-config`.
376. Write a program to patch a Secret’s data in the `test` namespace using `client-go`, updating `password` in `app-secret`.
377. Write a program to list all pods in the `test` namespace with label `app=nginx` using `client-go`, printing their names.
378. Write a program to list all services in the `test` namespace with annotation `managed=true` using `client-go`, printing their names.
379. Write a program to list all deployments in the `test` namespace with label `env=prod` using `client-go`, printing their replica counts.
380. Write a program to list all ConfigMaps in the `test` namespace with label `app=nginx` using `client-go`, printing their data keys.
381. Write a program to list all Secrets in the `test` namespace with annotation `secure=true` using `client-go`, printing their types.
382. Write a program to create a pod with a liveness probe in the `test` namespace using `client-go`, checking an `nginx` container’s health.
383. Write a program to create a pod with a readiness probe in the `test` namespace using `client-go`, ensuring an `nginx` container is ready.
384. Write a program to create a deployment with a rolling update strategy in the `test` namespace using `client-go`, deploying `nginx` containers.
385. Write a program to create a deployment with a recreate strategy in the `test` namespace using `client-go`, deploying `nginx` containers.
386. Write a program to create a service with session affinity in the `test` namespace using `client-go`, enabling sticky sessions for `nginx` pods.
387. Write a program to list all pods in the `test` namespace with label `app=nginx` using `client-go`, printing their status.
388. Write a program to create a ConfigMap with multiple keys in the `test` namespace using `client-go`, including `db=prod` and `env=test`.
389. Write a program to create a Secret with multiple keys in the `test` namespace using `client-go`, including `username` and `password`.
390. Write a program to create a pod with pod affinity in the `test` namespace using `client-go`, scheduling `nginx` pods near `redis` pods.
391. Write a program to create a deployment with resource requests in the `test` namespace using `client-go`, setting CPU and memory for `nginx` containers.
392. Write a program to patch a pod’s annotations in the `test` namespace using `client-go`, adding `managed=true` to a pod.
393. Write a program to patch a deployment’s labels in the `test` namespace using `client-go`, adding `app=core` to a deployment.
394. Write a program to patch a service’s port in the `test` namespace using `client-go`, changing `app-service` port to 8080.
395. Write a program to delete all pods with label `app=old` in the `test` namespace using `client-go`.
396. Write a program to delete all services with annotation `deprecated=true` in the `test` namespace using `client-go`.
397. Write a program to create a pod with a `hostPath` volume in the `test` namespace using `client-go`, mounting a local directory into an `nginx` container.
398. Write a program to create a pod with an init container in the `test` namespace using `client-go`, running a setup script before an `nginx` container.
399. Write a program to create a deployment with multiple containers in the `test` namespace using `client-go`, running `nginx` and `redis` in the same pod.
400. Write a program to create a service with multiple ports in the `test` namespace using `client-go`, exposing ports 80 and 443 for `nginx` pods.

---

### Notes
- **Assumptions**: Questions assume a Kubernetes cluster, a configured `client-go` client, and familiarity with Go. Resources like CRDs require registered schemas if referenced.
- **Variety**: Questions cover core Kubernetes resources (Pods, Services, Deployments, ConfigMaps, etc.), operations (list, create, update, delete, watch), and configurations (labels, annotations, volumes, probes).
- **Implementation**: Programs should handle errors, use proper contexts, and follow Kubernetes best practices (e.g., graceful deletion).
- **Dependencies**: Use `go mod` for `k8s.io/client-go` and related libraries.

This artifact provides 400 beginner-level questions to guide the development of Kubernetes-focused Go programs. For solutions, additional questions (intermediate/advanced), or specific implementations, please request further assistance.