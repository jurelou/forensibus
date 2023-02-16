/*
	This is a collection of rules that use external vriables
	They work with scanners that support the use of external variabls, like
	THOR, LOKI or SPARK
	https://www.nextron-systems.com/compare-our-scanners/
*/

rule test_forensibus {

	condition:
		filename == "test.exe" or filename == "AcroTray.exe"
}
