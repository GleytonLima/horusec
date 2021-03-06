// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package languages

import "strings"

type Language string

const (
	Go         Language = "Go"
	DotNet     Language = "C#"
	Ruby       Language = "Ruby"
	Python     Language = "Python"
	Java       Language = "Java"
	Kotlin     Language = "Kotlin"
	Javascript Language = "JavaScript"
	Leaks      Language = "Leaks"
	HCL        Language = "HCL"
	Unknown    Language = "Unknown"
)

func ParseStringToLanguage(value string) (l Language) {
	for key, lang := range l.MapEnableLanguages() {
		if strings.EqualFold(key, value) {
			return lang
		}
	}
	return Unknown
}

func SupportedLanguages() []Language {
	return []Language{
		Go,
		DotNet,
		Ruby,
		Python,
		Java,
		Kotlin,
		Javascript,
		Leaks,
		HCL,
		Unknown,
	}
}

func (l Language) MapEnableLanguages() map[string]Language {
	return map[string]Language{
		Go.ToString():         Go,
		Leaks.ToString():      Leaks,
		DotNet.ToString():     DotNet,
		Ruby.ToString():       Ruby,
		Python.ToString():     Python,
		Java.ToString():       Java,
		Kotlin.ToString():     Kotlin,
		Javascript.ToString(): Javascript,
		HCL.ToString():        HCL,
	}
}

func (l Language) ToString() string {
	return string(l)
}
