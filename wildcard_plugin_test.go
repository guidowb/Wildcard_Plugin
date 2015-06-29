//this file, wildcard_plugin_test.go, is created by '$ ginkgo generate wildcard_plugin.go'
//VOCAB:
//GetAppsStub = func() ([]plugin_models.GetAppsModel, error)
//
package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cloudfoundry/cli/plugin/models"
	"github.com/cloudfoundry/cli/plugin/fakes"
	// "github.com/cloudfoundry/cli/cf/configuration/core_config"
	//. "github.com/cloudfoundry/cli/testhelpers/matchers"
	//"github.com/onsi/gomega/matchers"
	// "fmt"
	// "reflect"
	//io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	testterm "github.com/cloudfoundry/cli/testhelpers/terminal"
	// testcmd "github.com/cloudfoundry/cli/testhelpers/commands"
	// testapi "github.com/cloudfoundry/cli/cf/api/fakes"






	. "github.com/cloudfoundry/cli/testhelpers/matchers"
)

//top-level describe container using Ginkgo's "Describe(text string, body func()) bool" cuntion.
//var_=.. allows us to eval Describe at the top level without the need to wrap it in "func init() {}"
var _ = Describe("WildcardPlugin", func() {

	Describe("Checking for correct results to getMatchedApps", func() {
		var (
			
			wildcardPlugin 		*Wildcard
			fakeCliConnection 	*fakes.FakeCliConnection
		)
		// runCommand := func(args ...string) bool {
		// 	wildcardPlugin.Run(fakeCliConnection, []string{"wc-a", "ca*"})
		// 	cmd := command_registry.Commands.FindCommand("apps")
		// 	return testcmd.RunCliCommand(cmd, args, requirementsFactory)
		// }
		appsList := make([]plugin_models.GetAppsModel, 0)
		appsList = append(appsList,
			plugin_models.GetAppsModel{"spring-music", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"spring-master", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"spring-nana", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"spring-spring", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"springtime", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"cake", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"carrot", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"car", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"c", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app1", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app2", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app3", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app4", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app5", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app10", "", "", 0, 0, 0, 0, nil},
			)

		BeforeEach(func() {
			
			fakeCliConnection = &fakes.FakeCliConnection{}
			wildcardPlugin = &Wildcard{}
		})

		Context("With wildcard asterisk(*)", func() {
			It("should return all apps starting with 'ca'", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				output := wildcardPlugin.getMatchedApps(fakeCliConnection, []string{"wc-a", "ca*"})
				Expect(len(output)).To(Equal(3))
				Expect(output[0].Name).To(Equal("cake"))
				Expect(output[1].Name).To(Equal("carrot"))
				Expect(output[2].Name).To(Equal("car"))
			})
			It("should return all apps", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				output := wildcardPlugin.getMatchedApps(fakeCliConnection, []string{"wc-a", "*"})
				Expect(len(output)).To(Equal(15))
				Expect(output[0].Name).To(Equal("spring-music"))
				Expect(output[1].Name).To(Equal("spring-master"))
				Expect(output[2].Name).To(Equal("spring-nana"))
				Expect(output[3].Name).To(Equal("spring-spring"))
				Expect(output[4].Name).To(Equal("springtime"))
				Expect(output[5].Name).To(Equal("cake"))
				Expect(output[6].Name).To(Equal("carrot"))
				Expect(output[7].Name).To(Equal("car"))
				Expect(output[8].Name).To(Equal("c"))
				Expect(output[9].Name).To(Equal("app1"))
				Expect(output[10].Name).To(Equal("app2"))
				Expect(output[11].Name).To(Equal("app3"))
				Expect(output[12].Name).To(Equal("app4"))
				Expect(output[13].Name).To(Equal("app5"))
				Expect(output[14].Name).To(Equal("app10"))		
			})
			It("should return all apps starting with 'sp'", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				output := wildcardPlugin.getMatchedApps(fakeCliConnection, []string{"wc-a", "sp*"})
				Expect(len(output)).To(Equal(5))
				Expect(output[0].Name).To(Equal("spring-music"))
				Expect(output[1].Name).To(Equal("spring-master"))
				Expect(output[2].Name).To(Equal("spring-nana"))
				Expect(output[3].Name).To(Equal("spring-spring"))
				Expect(output[4].Name).To(Equal("springtime"))				
			})
			It("should return all apps starting with 'app'", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				output := wildcardPlugin.getMatchedApps(fakeCliConnection, []string{"wc-a", "app*"})
				Expect(len(output)).To(Equal(6))
				Expect(output[0].Name).To(Equal("app1"))
				Expect(output[1].Name).To(Equal("app2"))
				Expect(output[2].Name).To(Equal("app3"))
				Expect(output[3].Name).To(Equal("app4"))
				Expect(output[4].Name).To(Equal("app5"))
				Expect(output[5].Name).To(Equal("app10"))
			})
		})
		Context("With wildcard question-mark(?)", func() {
			It("should return all apps with patter 'ca?'", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				output := wildcardPlugin.getMatchedApps(fakeCliConnection, []string{"wc-a", "ca?"})
				Expect(len(output)).To(Equal(1))
				Expect(output[0].Name).To(Equal("car"))
			})
			It("should return all apps with patter 'app?'", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				output := wildcardPlugin.getMatchedApps(fakeCliConnection, []string{"wc-a", "app?"})
				Expect(len(output)).To(Equal(5))
				Expect(output[0].Name).To(Equal("app1"))
				Expect(output[1].Name).To(Equal("app2"))
				Expect(output[2].Name).To(Equal("app3"))
				Expect(output[3].Name).To(Equal("app4"))
				Expect(output[4].Name).To(Equal("app5"))
			})
			It("should return all apps with patter 'app?'", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				output := wildcardPlugin.getMatchedApps(fakeCliConnection, []string{"wc-a", "app??"})
				Expect(len(output)).To(Equal(1))
				Expect(output[0].Name).To(Equal("app10"))
			})
		})
	})
	Describe("Checking for correct output from wildcard_plugin", func() {
		routeList := make([]plugin_models.GetAppsRouteSummary, 0)
		routeList = append(routeList,
			plugin_models.GetAppsRouteSummary{"1234", "www", plugin_models.GetAppsDomainFields{"1234", "google.com", "12345", false}},
			plugin_models.GetAppsRouteSummary{"5678", "www",  plugin_models.GetAppsDomainFields{"5678", "yahoo.com",  "12345", false}},
		)
		appsList := make([]plugin_models.GetAppsModel, 0)
		appsList = append(appsList,
			plugin_models.GetAppsModel{"spring-music", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"spring-master", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"spring-nana", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"spring-spring", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"springtime", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"cake", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"carrot", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"car", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"c", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app1", "1234-5678-90", "started", 3, 4, 512, 1024, routeList},
			plugin_models.GetAppsModel{"app2", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app3", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app4", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app5", "", "", 0, 0, 0, 0, nil},
			plugin_models.GetAppsModel{"app10", "", "", 0, 0, 0, 0, nil},
		)
		var (
			wildcardPlugin *Wildcard
			fakeCliConnection *fakes.FakeCliConnection
			ui                  *testterm.FakeUI
		//	configRepo          core_config.Repository
		//	appSummaryRepo      *testapi.FakeAppSummaryRepo
		)
		BeforeEach(func() {
			ui = &testterm.FakeUI{}
			fakeCliConnection = &fakes.FakeCliConnection{}
			wildcardPlugin = &Wildcard{ ui: ui, }
		})
// THIS ONE ACTUALLY PANICS - PLEASE FIX THE CODE FOR IT, THEN RE-ENABLE THE TEST
//		Context("With no arguments", func() {
//			It("displays usage", func() {
//				wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-apps"})
//				Expect(ui.Outputs).To(ContainSubstrings(
//					[]string{"Usage:", "cf", "wildcard-apps", "APP_NAME_WITH_WILDCARD"},
//				))
//			})
//		})
		Context("With too many arguments", func() {
			It("displays usage", func() {
				wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-apps", "app1", "app2"})
				Expect(ui.Outputs).To(ContainSubstrings(
					[]string{"Usage:", "cf", "wildcard-apps", "APP_NAME_WITH_WILDCARD"},
				))
			})
		})
		Context("With wildcard asterisk(*)", func() {
			It("displays all columns", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-apps", "app1"})
				Expect(ui.Outputs).To(ContainSubstrings(
					[]string{"name", "requested state", "instances", "memory", "disk", "urls"},
					[]string{"app1", "started", "3/4", "512M", "1G", "www.google.com", "www.yahoo.com"},
				))
			})
			It("lists all apps", func() {
				fakeCliConnection.GetAppsReturns(appsList, nil)
				wildcardPlugin.Run(fakeCliConnection, []string{"wildcard-apps", "app*"})
				Expect(ui.Outputs).To(ContainSubstrings(
					[]string{"app1"},
					[]string{"app2"},
					[]string{"app3"},
					[]string{"app4"},
					[]string{"app5"},
				))
			})
		})
	})
})
