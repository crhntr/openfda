## Data copied from EXCEL
# total	other	lifeThreatening	hospitalization	disabling	death	congenitalAnomali
# total	1794	663	137	798	104	146	7
# male	955	283	81	520	48	65	6
# female	839	380	56	278	56	81	1
#
# 0	total	other	lifeThreatening	hospitalization	disabling	death	congenitalAnomali
# male	0.532329989	0.426847662	0.591240876	0.651629073	0.461538462	0.445205479	0.857142857
# female	0.467670011	0.573152338	0.408759124	0.348370927	0.538461538	0.554794521	0.142857143
#
# average prop.		0.369565217	0.076365663	0.444816054	0.057971014	0.081382386	0.003901895
# observed statistic		0.146304676	-0.182481752	-0.303258145	0.076923077	0.109589041	-0.714285714
# p-value			0	0	0	0	0

simulationNum = 1000

# lifeThreatening
pi = 137/1794
list_M = rbinom(simulationNum, 81, pi)/simulationNum
list_F = rbinom(simulationNum, 56, pi)/simulationNum
difs = list_M - list_F
pvalue = sum(difs <= -0.182481752)/simulationNum + sum(difs >= 0.182481752)/simulationNum
pvalue

# hospitalization
pi = 798/1794
list_M = rbinom(simulationNum, 520, pi)/simulationNum
list_F = rbinom(simulationNum, 278, pi)/simulationNum
difs = list_M - list_F
pvalue = sum(difs<=-0.303258145)/simulationNum + sum(difs>=0.303258145)/simulationNum
pvalue

# disabling
pi = 104/1794
list_M = rbinom(simulationNum, 48, pi)/simulationNum
list_F = rbinom(simulationNum, 56, pi)/simulationNum
difs = list_M - list_F
pvalue = sum(difs<=-0.076923077)/simulationNum + sum(difs>=0.076923077)/simulationNum
pvalue

# death
pi = 146/1794
list_M = rbinom(simulationNum, 65, pi)/simulationNum
list_F = rbinom(simulationNum, 81, pi)/simulationNum
difs = list_M - list_F
pvalue = sum(difs<=-0.109589041)/simulationNum + sum(difs>=0.109589041)/simulationNum
pvalue

# congenitalAnomali
pi = 7/1794
list_M = rbinom(simulationNum, 6, pi)/simulationNum
list_F = rbinom(simulationNum, 1, pi)/simulationNum
difs = list_M - list_F
pvalue = sum(difs<=-0.714285714)/simulationNum + sum(difs>=0.714285714)/simulationNum
pvalue
