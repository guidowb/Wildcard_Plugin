//this file, wildcard_plugin_test.go, is created by '$ ginkgo generate wildcard_plugin.go'
//VOCAB:
//GetAppsStub = func() ([]plugin_models.ApplicationSummary, error)
//
package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cloudfoundry/cli/plugin/models"
	"github.com/cloudfoundry/cli/plugin/fakes"
	"fmt"
	"reflect"
)

//top-level describe container using Ginkgo's "Describe(text string, body func()) bool" cuntion.
//var_=.. allows us to eval Describe at the top level without the need to wrap it in "func init() {}"
var _ = Describe("WildcardPlugin", func() {

	Describe("Checking for correct results to getMatchedApps", func() {
		var (
			wildcardPlugin Wildcard
			fakeCliConnection *fakes.FakeCliConnection
		)
		appsList := make([]plugin_models.ApplicationSummary, 0)
		appsList = append(appsList,
			plugin_models.ApplicationSummary{"spring-music", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"spring-master", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"spring-nana", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"spring-spring", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"springtime", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"cake", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"carrot", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"car", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"c", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app1", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app2", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app3", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app4", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app5", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app10", "", "", 0, 0, 0, 0, nil},
			)
		BeforeEach(func() {
			fakeCliConnection = &fakes.FakeCliConnection{}
			wildcardPlugin = Wildcard{}
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
	Describe("Checking for correct results to WildcardCommandApps", func() {
		var (
			wildcardPlugin Wildcard
			fakeCliConnection *fakes.FakeCliConnection
		)
		appsList := make([]plugin_models.ApplicationSummary, 0)
		appsList = append(appsList,
			plugin_models.ApplicationSummary{"spring-music", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"spring-master", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"spring-nana", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"spring-spring", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"springtime", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"cake", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"carrot", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"car", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"c", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app1", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app2", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app3", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app4", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app5", "", "", 0, 0, 0, 0, nil},
			plugin_models.ApplicationSummary{"app10", "", "", 0, 0, 0, 0, nil},
		)
		BeforeEach(func() {
			fakeCliConnection = &fakes.FakeCliConnection{}
			wildcardPlugin = Wildcard{}
		})
		Context("With wildcard asterisk(*)", func() {
			It("should return all apps starting with 'ca'", func() {
				output, _ := fakeCliConnection.CliCommandWithoutTerminalOutput("wc-a", "ca*")

				//fmt.Println(output)
				//fmt.Println(output[0])
				//fmt.Println(output[1])
				
				fmt.Println(reflect.TypeOf(output))
				fmt.Println(output[1][0])
				fmt.Println(output[1][1])
				// for idx, v := range output {
				// 	v = strings.TrimSpace(v)
				// 	if strings.HasPrefix(v, "FAILED") {
				// 		e := output[idx+1]
				// 		return status, errors.New(e)
				// 	}
				// 	if strings.HasPrefix(v, "requested state: ") {
				// 		status.state = strings.TrimPrefix(v, "requested state: ")
				// 	}
				// 	if strings.HasPrefix(v, "instances: ") {
				// 		instances := strings.TrimPrefix(v, "instances: ")
				// 		split := strings.Split(instances, "/")
				// 		status.countRunning, _ = strconv.Atoi(split[0])
				// 		status.countRequested, _ = strconv.Atoi(split[1])
				// 	}
				// 	if strings.HasPrefix(v, "urls: ") {
				// 		urls := strings.TrimPrefix(v, "urls: ")
				// 		status.routes = strings.Split(urls, ", ")
				// 	}
				// }

				// Expect(len(output)).To(Equal(3))
				// Expect(output[0].Name).To(Equal("cake"))
				// Expect(output[1].Name).To(Equal("carrot"))
				// Expect(output[2].Name).To(Equal("car"))


			})

			// It("should return all apps starting with 'sp'", func() {
			// 	fakeCliConnection.GetAppsReturns(appsList, nil)
			// 	var err error
			// 	_, err = wildcardPlugin.getAppStatus(fakeCliConnection, "app1")
			// 	Expect(err.Error()).To(Equal("App app1 not found"))
			// })
		})
	// 	Context("With wildcard sp*", func() {
	// 		It("should return all apps starting with 'sp'", func() {
	// 			fakeCliConnection.GetAppsReturns(appsList, nil)
	// 			var err error
	// 			_, err = wildcardPlugin.getAppStatus(fakeCliConnection, "app1")
	// 			Expect(err.Error()).To(Equal("App app1 not found"))
	// 		})
	// 		It("should return all apps starting with 'sp'", func() {
	// 			fakeCliConnection.GetAppsReturns(appsList, nil)
	// 			var err error
	// 			_, err = wildcardPlugin.getAppStatus(fakeCliConnection, "app1")
	// 			Expect(err.Error()).To(Equal("App app1 not found"))
	// 		})
	// 	})
	// })
	// BeforeEach(func() {
	// 	fakeCliConnection = &fakes.FakeCliConnection{}
	// 	scaleoverCmdPlugin = &ScaleoverCmd{}
	})

})
