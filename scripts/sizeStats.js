// var c = db.drug_event.find({}, {"dgs.openfda.bn": 1, "dgs.openfda.gn": 1, "dgs.openfda.pt": "HUMAN PRESCRIPTION DRUG"}).limit(10000)
var maxBN = "", maxGN = "", sizeBN = 0, sizeGN = 0, countBN = 0, countGN = 0;

while (c.hasNext()) {
  var doc = c.next();
  if (doc && doc.dgs) {
    for (var di in doc.dgs) {
      var drg = doc.dgs[di];

      if (drg.openfda) {
        if (drg.openfda.bn) {
          var bns = drg.openfda.bn;
          for (var i = 0; i < bns.length; i++) {
            var bn = bns[i];
            if (bn.length > maxBN.length) { maxBN = bn }
            sizeBN += bn.length;
            countBN++;
          }
        }

        if (drg.openfda.gn) {
          var gns = drg.openfda.gn;
          for (var i = 0; i < gns.length; i++) {
            var gn = gns[i];
            if (gn.length > maxGN.length) { maxGN = gn }
            sizeGN += gn.length;
            countGN++;
          }
        }
      }
    }
  }
}

printjson({
  maxBN, maxGN, sizeBN, sizeGN, count,
  avgBN: sizeBN/countBN,
  avgGN: sizeGN/countGN,
})
