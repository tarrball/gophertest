package expects

func (source ExpectationBool) toBeFalse() {
	if source.value != false {
		source.testContext.Errorf("Expected 'true' to be 'false'")
	}
}
