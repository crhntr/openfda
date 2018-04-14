package drug

import "testing"

func TestEventBSONKeys(t *testing.T) {
	m := make(map[string]int)
	for i, key := range []string{BSONKeyEventFileName,
		BSONKeyEventFileHash,
		BSONKeyEventSafetyReportID,
		BSONKeyEventSafetyReportVersion,
		BSONKeyEventReceiveDate,
		BSONKeyEventReceiptDate,
		BSONKeyEventTransmissionDate,
		BSONKeyEventSeriousness,
		BSONKeyEventDuplicate,
		BSONKeyEventCompanyNumber,
		BSONKeyEventOccurCountry,
		BSONKeyEventPrimarySourceQualification,
		BSONKeyEventPrimarySourceReporterCountry,
		BSONKeyEventReportDuplicateSource,
		BSONKeyEventReportDuplicateNumber,
		BSONKeyEventSenderType,
		BSONKeyEventSenderOrganization,
		BSONKeyEventReceiverType,
		BSONKeyEventReceiverOrganization,
		BSONKeyEventOnsetAge,
		BSONKeyEventOnsetAgeUnit,
		BSONKeyEventSex,
		BSONKeyEventWeight,
		BSONKeyEventDeathDate,
		BSONKeyEventDrugs,
		BSONKeyEventActionDrug,
		BSONKeyEventAdditional,
		BSONKeyEventCumulativeDosageNumber,
		BSONKeyEventCumulativeDosageUnit,
		BSONKeyEventDosageForm,
		BSONKeyEventIntervalDosageDefinition,
		BSONKeyEventIntervalDosageUnitNumber,
		BSONKeyEventRecurreAdministration,
		BSONKeyEventSeparateDosageNumber,
		BSONKeyEventStructureDosageNumber,
		BSONKeyEventStructureDosageUnit,
		BSONKeyEventAdministrationRoute,
		BSONKeyEventAuthorizationNumber,
		BSONKeyEventBatchNumber,
		BSONKeyEventCharacterization,
		BSONKeyEventDoseagetext,
		BSONKeyEventStartDate,
		BSONKeyEventEndDate,
		BSONKeyEventIndication,
		BSONKeyEventTreatmentDuration,
		BSONKeyEventOpenFDA,
		BSONKeyEventReactions,
		BSONKeyEventMedDRA,
		BSONKeyEventMedDRAVersion,
		BSONKeyEventOutcome} {
		if _, ok := m[key]; ok {
			t.Errorf("duplicate key at %d %q", i, key)
		}
		m[key] += 1
	}
}
