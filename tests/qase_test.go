/*
Copyright © 2023 - 2024 SUSE LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package qase_test

import (
	"fmt"
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/rancher-sandbox/qase-ginkgo"
)

// Just a dumb function for testing purposes
func outputCurrentPwd() ([]byte, error) {
	return (exec.Command("pwd").Output())
}

var _ = Describe("Qase Ginkgo Integration - Unit tests", func() {
	var qaseRunID int32

	It("Create and export a Qase Run", func() {
		qaseRunID = CreateRun()
		GinkgoWriter.Printf("Run ID %d created\n", qaseRunID)
		Expect(qaseRunID).To(BeNumerically(">", 0))
	})

	// The whole test with ID=1 should be marked as failed
	It("Test the Qase function with ID=1", func() {
		// Report to Qase
		testCaseID = 1

		// TODO: could be a good idea to add a function to check a case status
		By("testing that output is not empty (will pass)", func() {
			Ω(outputCurrentPwd()).Should(Not(BeEmpty()))
		})
	})

	// Finalize and report the result
	It("Finalize the run and publish the Qase report", func() {
		By("finalizing and reporting", func() {
			url := FinalizeResults()
			GinkgoWriter.Printf("Results URL: %s\n", url)
			Expect(url).To(Not(BeEmpty()))
		})
	})

	It("Delete the previously generated Qase run", func() {
		By("Deleting run id "+fmt.Sprint(qaseRunID), func() {
			// No real check here, as the function is supposed to trigger a Fatal call in case of issue
			DeleteRun()
		})
	})
})
