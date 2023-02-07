# NebulaGraph Query Compare

```
# build
make

# sample 1001 samples per 200 records
./compare_query sample --source ~/person.csv --sample person.csv

# compare the queries, would compare queries in ./pkg/tpl/
# would print the first query for debugging
./compare_query compare --sample ./person.csv --nebulagraph 192.168.15.8:9669 --space sf100

```

result
```
Now(): 2023-02-07 17:36:39.73277055 +0800 CST m=+0.005404806, run 0 iterations. 
left query: 
GO FROM 933 OVER KNOWS YIELD KNOWS._dst as one_hop_id, timestamp(KNOWS.creationDate) as one_hop_edge_strength | ORDER BY $-.one_hop_edge_strength DESC | LIMIT 20 | GO FROM $-.one_hop_id OVER KNOWS YIELD DISTINCT KNOWS._dst as two_hop_id, timestamp(KNOWS.creationDate) + $-.one_hop_edge_strength as two_hop_edge_strength | ORDER BY $-.two_hop_edge_strength DESC | LIMIT 100 | GO FROM $-.two_hop_id OVER KNOWS YIELD KNOWS._dst as three_hop_id, timestamp(KNOWS.creationDate) + $-.two_hop_edge_strength as three_hop_edge_strength | ORDER BY $-.three_hop_edge_strength DESC | LIMIT 2000 | YIELD DISTINCT $-.three_hop_id as three_hop_id; 
right query: 
match (v)-[e1:KNOWS]->(v1) where id(v) == 933 with distinct v1, timestamp(e1.creationDate) as one_hop_edge_strength order by one_hop_edge_strength desc limit 20 match (v1)-[e2:KNOWS]->(v2) with distinct v1, one_hop_edge_strength, v2, timestamp(e2.creationDate) as tmp_two_hop_edge_strength with distinct v2,tmp_two_hop_edge_strength+one_hop_edge_strength as two_hop_edge_strength order by two_hop_edge_strength desc limit 100 match (v2)-[e3:KNOWS]->(v3) with distinct v2, two_hop_edge_strength, v3,  timestamp(e3.creationDate) as tmp_three_hop_edge_strength with distinct v3, tmp_three_hop_edge_strength+two_hop_edge_strength as three_hop_edge_strength order by three_hop_edge_strength desc limit 2000 return DISTINCT id(v3) as three_hop_id 

left query: 
GO FROM 933 OVER KNOWS YIELD KNOWS._dst as one_hop_id, timestamp(KNOWS.creationDate) as one_hop_edge_strength | ORDER BY $-.one_hop_edge_strength DESC | LIMIT 20 | GO FROM $-.one_hop_id OVER KNOWS YIELD DISTINCT KNOWS._dst as two_hop_id, timestamp(KNOWS.creationDate) + $-.one_hop_edge_strength as two_hop_edge_strength | ORDER BY $-.two_hop_edge_strength DESC | LIMIT 100 | GO FROM $-.two_hop_id OVER KNOWS YIELD KNOWS._dst as three_hop_id, timestamp(KNOWS.creationDate) + $-.two_hop_edge_strength as three_hop_edge_strength | ORDER BY $-.three_hop_edge_strength DESC | LIMIT 2000 | YIELD DISTINCT $-.three_hop_id as three_hop_id; 
right query: 
match (v)-[e1:KNOWS]->(v1) where id(v) == 933 with distinct v1, timestamp(e1.creationDate) as one_hop_edge_strength order by one_hop_edge_strength desc limit 20 match (v1)-[e2:KNOWS]->(v2) with distinct v1, one_hop_edge_strength, v2, timestamp(e2.creationDate) as tmp_two_hop_edge_strength with distinct v2,tmp_two_hop_edge_strength+one_hop_edge_strength as two_hop_edge_strength order by two_hop_edge_strength desc limit 100 match (v2)-[e3:KNOWS]->(v3) with distinct v2, two_hop_edge_strength, v3,  timestamp(e3.creationDate) as tmp_three_hop_edge_strength with distinct v3, tmp_three_hop_edge_strength+two_hop_edge_strength as three_hop_edge_strength order by three_hop_edge_strength desc limit 2000 return DISTINCT id(v3) as three_hop_id 

Now(): 2023-02-07 17:36:40.030870449 +0800 CST m=+0.303504705, run 10 iterations. 
Now(): 2023-02-07 17:36:40.314908825 +0800 CST m=+0.587543080, run 20 iterations. 
Now(): 2023-02-07 17:36:40.628243263 +0800 CST m=+0.900877520, run 30 iterations. 
Now(): 2023-02-07 17:36:40.981538656 +0800 CST m=+1.254172912, run 40 iterations. 
Now(): 2023-02-07 17:36:41.234524059 +0800 CST m=+1.507158315, run 50 iterations. 
Now(): 2023-02-07 17:36:41.639493847 +0800 CST m=+1.912128103, run 60 iterations. 
Now(): 2023-02-07 17:36:41.976437544 +0800 CST m=+2.249071800, run 70 iterations. 
Now(): 2023-02-07 17:36:42.367821526 +0800 CST m=+2.640455791, run 80 iterations. 
Now(): 2023-02-07 17:36:42.725598558 +0800 CST m=+2.998232814, run 90 iterations. 
Now(): 2023-02-07 17:36:43.02555476 +0800 CST m=+3.298189016, run 100 iterations. 
Now(): 2023-02-07 17:36:43.370133048 +0800 CST m=+3.642767304, run 110 iterations. 
Now(): 2023-02-07 17:36:43.653116543 +0800 CST m=+3.925750798, run 120 iterations. 
Now(): 2023-02-07 17:36:43.964218764 +0800 CST m=+4.236853020, run 130 iterations. 
Now(): 2023-02-07 17:36:44.289582089 +0800 CST m=+4.562216345, run 140 iterations. 
Now(): 2023-02-07 17:36:44.632829233 +0800 CST m=+4.905463490, run 150 iterations. 
Now(): 2023-02-07 17:36:44.90973504 +0800 CST m=+5.182369304, run 160 iterations. 
Now(): 2023-02-07 17:36:45.231553286 +0800 CST m=+5.504187543, run 170 iterations. 
Now(): 2023-02-07 17:36:45.606011441 +0800 CST m=+5.878645697, run 180 iterations. 
Now(): 2023-02-07 17:36:45.907408885 +0800 CST m=+6.180043142, run 190 iterations. 
Now(): 2023-02-07 17:36:46.211584045 +0800 CST m=+6.484218301, run 200 iterations. 
Now(): 2023-02-07 17:36:46.435222364 +0800 CST m=+6.707856620, run 210 iterations. 
Now(): 2023-02-07 17:36:46.763172473 +0800 CST m=+7.035806729, run 220 iterations. 
Now(): 2023-02-07 17:36:47.143495146 +0800 CST m=+7.416129403, run 230 iterations. 
Now(): 2023-02-07 17:36:47.503219938 +0800 CST m=+7.775854194, run 240 iterations. 
Now(): 2023-02-07 17:36:47.824025888 +0800 CST m=+8.096660145, run 250 iterations. 
Now(): 2023-02-07 17:36:48.111951199 +0800 CST m=+8.384585456, run 260 iterations. 
Now(): 2023-02-07 17:36:48.478367662 +0800 CST m=+8.751001927, run 270 iterations. 
Now(): 2023-02-07 17:36:48.736071938 +0800 CST m=+9.008706194, run 280 iterations. 
Now(): 2023-02-07 17:36:49.094683646 +0800 CST m=+9.367317903, run 290 iterations. 
Now(): 2023-02-07 17:36:49.448564329 +0800 CST m=+9.721198599, run 300 iterations. 
Now(): 2023-02-07 17:36:49.767480005 +0800 CST m=+10.040114261, run 310 iterations. 
Now(): 2023-02-07 17:36:50.100927033 +0800 CST m=+10.373561289, run 320 iterations. 
Now(): 2023-02-07 17:36:50.41806974 +0800 CST m=+10.690703996, run 330 iterations. 
Now(): 2023-02-07 17:36:50.749894576 +0800 CST m=+11.022528832, run 340 iterations. 
Now(): 2023-02-07 17:36:51.060376799 +0800 CST m=+11.333011054, run 350 iterations. 
Now(): 2023-02-07 17:36:51.376887133 +0800 CST m=+11.649521389, run 360 iterations. 
Now(): 2023-02-07 17:36:51.696132612 +0800 CST m=+11.968766868, run 370 iterations. 
Now(): 2023-02-07 17:36:52.031630044 +0800 CST m=+12.304264300, run 380 iterations. 
Now(): 2023-02-07 17:36:52.378151028 +0800 CST m=+12.650785284, run 390 iterations. 
Now(): 2023-02-07 17:36:52.707268398 +0800 CST m=+12.979902654, run 400 iterations. 
Now(): 2023-02-07 17:36:53.07524452 +0800 CST m=+13.347878776, run 410 iterations. 
Now(): 2023-02-07 17:36:53.324770929 +0800 CST m=+13.597405185, run 420 iterations. 
Now(): 2023-02-07 17:36:53.643856799 +0800 CST m=+13.916491055, run 430 iterations. 
Now(): 2023-02-07 17:36:53.967160711 +0800 CST m=+14.239794967, run 440 iterations. 
Now(): 2023-02-07 17:36:54.357757887 +0800 CST m=+14.630392144, run 450 iterations. 
Now(): 2023-02-07 17:36:54.703836413 +0800 CST m=+14.976470687, run 460 iterations. 
Now(): 2023-02-07 17:36:54.98063696 +0800 CST m=+15.253271216, run 470 iterations. 
Now(): 2023-02-07 17:36:55.315355252 +0800 CST m=+15.587989526, run 480 iterations. 
Now(): 2023-02-07 17:36:55.557175515 +0800 CST m=+15.829809771, run 490 iterations. 
Now(): 2023-02-07 17:36:55.868171622 +0800 CST m=+16.140805878, run 500 iterations. 
Now(): 2023-02-07 17:36:56.14550619 +0800 CST m=+16.418140446, run 510 iterations. 
Now(): 2023-02-07 17:36:56.41437836 +0800 CST m=+16.687012625, run 520 iterations. 
Now(): 2023-02-07 17:36:56.678514791 +0800 CST m=+16.951149047, run 530 iterations. 
Now(): 2023-02-07 17:36:56.95603777 +0800 CST m=+17.228672026, run 540 iterations. 
Now(): 2023-02-07 17:36:57.349127814 +0800 CST m=+17.621762070, run 550 iterations. 
Now(): 2023-02-07 17:36:57.675231522 +0800 CST m=+17.947865779, run 560 iterations. 
Now(): 2023-02-07 17:36:57.948386655 +0800 CST m=+18.221020910, run 570 iterations. 
Now(): 2023-02-07 17:36:58.254711721 +0800 CST m=+18.527345987, run 580 iterations. 
Now(): 2023-02-07 17:36:58.542345232 +0800 CST m=+18.814979488, run 590 iterations. 
Now(): 2023-02-07 17:36:58.900996973 +0800 CST m=+19.173631229, run 600 iterations. 
Now(): 2023-02-07 17:36:59.281940748 +0800 CST m=+19.554575004, run 610 iterations. 
Now(): 2023-02-07 17:36:59.590847248 +0800 CST m=+19.863481504, run 620 iterations. 
Now(): 2023-02-07 17:36:59.970822539 +0800 CST m=+20.243456795, run 630 iterations. 
Now(): 2023-02-07 17:37:00.243958538 +0800 CST m=+20.516592793, run 640 iterations. 
Now(): 2023-02-07 17:37:00.58937406 +0800 CST m=+20.862008316, run 650 iterations. 
Now(): 2023-02-07 17:37:00.865082097 +0800 CST m=+21.137716353, run 660 iterations. 
Now(): 2023-02-07 17:37:01.147611939 +0800 CST m=+21.420246196, run 670 iterations. 
Now(): 2023-02-07 17:37:01.475715372 +0800 CST m=+21.748349629, run 680 iterations. 
Now(): 2023-02-07 17:37:01.805240869 +0800 CST m=+22.077875125, run 690 iterations. 
Now(): 2023-02-07 17:37:02.130135913 +0800 CST m=+22.402770171, run 700 iterations. 
Now(): 2023-02-07 17:37:02.433821012 +0800 CST m=+22.706455268, run 710 iterations. 
Now(): 2023-02-07 17:37:02.745386767 +0800 CST m=+23.018021024, run 720 iterations. 
Now(): 2023-02-07 17:37:03.020960355 +0800 CST m=+23.293594611, run 730 iterations. 
Now(): 2023-02-07 17:37:03.359711632 +0800 CST m=+23.632345888, run 740 iterations. 
Now(): 2023-02-07 17:37:03.677088886 +0800 CST m=+23.949723143, run 750 iterations. 
Now(): 2023-02-07 17:37:03.974551304 +0800 CST m=+24.247185563, run 760 iterations. 
Now(): 2023-02-07 17:37:04.250741167 +0800 CST m=+24.523375423, run 770 iterations. 
Now(): 2023-02-07 17:37:04.518333588 +0800 CST m=+24.790967844, run 780 iterations. 
Now(): 2023-02-07 17:37:04.835438707 +0800 CST m=+25.108072963, run 790 iterations. 
Now(): 2023-02-07 17:37:05.180866078 +0800 CST m=+25.453500333, run 800 iterations. 
Now(): 2023-02-07 17:37:05.425330551 +0800 CST m=+25.697964808, run 810 iterations. 
Now(): 2023-02-07 17:37:05.742415634 +0800 CST m=+26.015049891, run 820 iterations. 
Now(): 2023-02-07 17:37:05.999111483 +0800 CST m=+26.271745739, run 830 iterations. 
Now(): 2023-02-07 17:37:06.300251576 +0800 CST m=+26.572885832, run 840 iterations. 
Now(): 2023-02-07 17:37:06.603495416 +0800 CST m=+26.876129672, run 850 iterations. 
Now(): 2023-02-07 17:37:06.933969569 +0800 CST m=+27.206603825, run 860 iterations. 
Now(): 2023-02-07 17:37:07.248672724 +0800 CST m=+27.521306980, run 870 iterations. 
Now(): 2023-02-07 17:37:07.529995566 +0800 CST m=+27.802629822, run 880 iterations. 
Now(): 2023-02-07 17:37:07.811989349 +0800 CST m=+28.084623606, run 890 iterations. 
Now(): 2023-02-07 17:37:08.097206977 +0800 CST m=+28.369841233, run 900 iterations. 
Now(): 2023-02-07 17:37:08.446602874 +0800 CST m=+28.719237139, run 910 iterations. 
Now(): 2023-02-07 17:37:08.732410638 +0800 CST m=+29.005044895, run 920 iterations. 
Now(): 2023-02-07 17:37:09.028603543 +0800 CST m=+29.301237799, run 930 iterations. 
Now(): 2023-02-07 17:37:09.303881814 +0800 CST m=+29.576516079, run 940 iterations. 
Now(): 2023-02-07 17:37:09.606940012 +0800 CST m=+29.879574268, run 950 iterations. 
Now(): 2023-02-07 17:37:09.889488164 +0800 CST m=+30.162122421, run 960 iterations. 
Now(): 2023-02-07 17:37:10.149807558 +0800 CST m=+30.422441814, run 970 iterations. 
Now(): 2023-02-07 17:37:10.467418076 +0800 CST m=+30.740052331, run 980 iterations. 
Now(): 2023-02-07 17:37:10.73767409 +0800 CST m=+31.010308346, run 990 iterations. 
Now(): 2023-02-07 17:37:11.058612884 +0800 CST m=+31.331247140, run 1000 iterations. 

```