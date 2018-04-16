# total	other	lifeThreatening	hospitalization	disabling	death	congenitalAnomali
# total	1794	663	137	798	104	146	7
# male	955	283	81	520	48	65	6
# female	839	380	56	278	56	81	1
#
#
#
# 0	total	other	lifeThreatening	hospitalization	disabling	death	congenitalAnomali
# male	0.532329989	0.296335079	0.084816754	0.544502618	0.05026178	0.068062827	0.006282723
# female	0.467670011	0.452920143	0.066746126	0.331346841	0.066746126	0.096543504	0.001191895
#
# average prob.		0.369565217	0.076365663	0.444816054	0.057971014	0.081382386	0.003901895
# observed statistic		0.156585064	-0.018070628	-0.213155776	0.016484346	0.028480677	-0.005090827

simulationNum = 100000

n_m = 955
n_f = 839

# lifeThreatening
pi = 137/1794
observedStatistic = -0.018070628
list_m = rbinom(n=simulationNum, n_m, prob=pi)/n_m
list_f = rbinom(n=simulationNum, n_f, prob=pi)/n_f
difs = list_m - list_f
pvalueLeft = sum(difs<=-observedStatistic)/simulationNum
pvalueRight = sum(difs>=observedStatistic)/simulationNum
pvalueSum = pvalueLeft + pvalueRight
pvalueRight
pvalueLeft
pvalueSum

# hospitalization
pi = 798/1794
observedStatistic = -0.213155776
list_m = rbinom(n=simulationNum, n_m, prob=pi)/n_m
list_f = rbinom(n=simulationNum, n_f, prob=pi)/n_f
difs = list_m - list_f
pvalueLeft = sum(difs<=-observedStatistic)/simulationNum
pvalueRight = sum(difs>=observedStatistic)/simulationNum
pvalueSum = pvalueLeft + pvalueRight
pvalueRight
pvalueLeft
pvalueSum

# disabling
pi = 104/1794
observedStatistic = 0.016484346
list_m = rbinom(n=simulationNum, n_m, prob=pi)/n_m
list_f = rbinom(n=simulationNum, n_f, prob=pi)/n_f
difs = list_m - list_f
pvalueLeft = sum(difs<=-observedStatistic)/simulationNum
pvalueRight = sum(difs>=observedStatistic)/simulationNum
pvalueSum = pvalueLeft + pvalueRight
pvalueRight
pvalueLeft
pvalueSum

# death
pi = 146/1794
observedStatistic = 0.028480677
list_m = rbinom(n=simulationNum, n_m, prob=pi)/n_m
list_f = rbinom(n=simulationNum, n_f, prob=pi)/n_f
difs = list_m - list_f
pvalueLeft = sum(difs<=-observedStatistic)/simulationNum
pvalueRight = sum(difs>=observedStatistic)/simulationNum
pvalueSum = pvalueLeft + pvalueRight
pvalueRight
pvalueLeft
pvalueSum

# congenitalAnomali
pi = 7/1794
observedStatistic = -0.005090827
list_m = rbinom(n=simulationNum, n_m, prob=pi)/n_m
list_f = rbinom(n=simulationNum, n_f, prob=pi)/n_f
difs = list_m - list_f
pvalueLeft = sum(difs<=-observedStatistic)/simulationNum
pvalueRight = sum(difs>=observedStatistic)/simulationNum
pvalueSum = pvalueLeft + pvalueRight
pvalueRight
pvalueLeft
pvalueSum
