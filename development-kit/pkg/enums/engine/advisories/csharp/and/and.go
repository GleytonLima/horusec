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

//nolint:lll multiple regex is not possible broken lines
package and

import (
	engine "github.com/ZupIT/horusec-engine"
	"github.com/ZupIT/horusec-engine/text"
	"github.com/ZupIT/horusec/development-kit/pkg/enums/confidence"
	"github.com/ZupIT/horusec/development-kit/pkg/enums/severity"
	"regexp"
)

func NewCsharpAndCommandInjection() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "782ad071-1cf3-4230-936f-b7a1e794828d",
			Name:        "Command Injection",
			Description: "If a malicious user controls either the FileName or Arguments, he might be able to execute unwanted commands or add unwanted argument. This behavior would not be possible if input parameter are validate against a white-list of characters. For more information access: (https://security-code-scan.github.io/#SCS0001).",
			Severity:    severity.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`new Process\(\)`),
			regexp.MustCompile(`StartInfo.FileName`),
			regexp.MustCompile(`StartInfo.Arguments`),
		},
	}
}

func NewCsharpAndXPathInjection() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "010a99b2-9f35-4cb3-9209-248bacba07f8",
			Name:        "XPath Injection",
			Description: "If the user input is not properly filtered, a malicious user could extend the XPath query. For more information access: (https://security-code-scan.github.io/#SCS0003).",
			Severity:    severity.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`new XmlDocument {XmlResolver = null}`),
			regexp.MustCompile(`Load\(.*\)`),
			regexp.MustCompile(`SelectNodes\(.*\)`),
		},
	}
}

func NewCsharpAndExternalEntityInjection() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "f3390c7e-8151-4f8a-8bc4-433841569153",
			Name:        "XML eXternal Entity Injection (XXE)",
			Description: "The XML parser is configured incorrectly. The operation could be vulnerable to XML eXternal Entity (XXE) processing. For more information access: (https://security-code-scan.github.io/#SCS0007).",
			Severity:    severity.Medium.ToString(),
			Confidence:  confidence.High.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`new XmlReaderSettings\(\)`),
			regexp.MustCompile(`XmlReader.Create\(.*\)`),
			regexp.MustCompile(`new XmlDocument\(.*\)`),
			regexp.MustCompile(`Load\(.*\)`),
			regexp.MustCompile(`ProhibitDtd = false`),
			regexp.MustCompile(`(new XmlReaderSettings\(\))(([^P]|P[^r]|Pr[^o]|Pro[^h]|Proh[^i]|Prohi[^b]|Prohib[^i]|Prohibi[^t]|Prohibit[^D]|ProhibitD[^t]|ProhibitDt[^d])*)(\.Load\(.*\))`),
		},
	}
}

func NewCsharpAndPathTraversal() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "3b905eb5-5af7-41db-a234-38c934f675b2",
			Name:        "Path Traversal",
			Description: "A path traversal attack (also known as directory traversal) aims to access files and directories that are stored outside the expected directory.By manipulating variables that reference files with “dot-dot-slash (../)” sequences and its variations or by using absolute file paths, it may be possible to access arbitrary files and directories stored on file system including application source code or configuration and critical system files. For more information access: (https://security-code-scan.github.io/#SCS0018).",
			Severity:    severity.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`ActionResult`),
			regexp.MustCompile(`System.IO.File.ReadAllBytes\(Server.MapPath\(.*\) \+ .*\)`),
			regexp.MustCompile(`File\(.*, System.Net.Mime.MediaTypeNames.Application.Octet, .*\)`),
		},
	}
}

func NewCsharpAndSQLInjectionWebControls() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "b6f82b7c-f321-4651-8ad3-87fbf5e0412b",
			Name:        "SQL Injection WebControls",
			Description: "Malicious user might get direct read and/or write access to the database. If the database is poorly configured the attacker might even get Remote Code Execution (RCE) on the machine running the database. For more information access: (https://security-code-scan.github.io/#SCS0014).",
			Severity:    severity.High.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`"Select .* From .* where .*" & .*`),
			regexp.MustCompile(`System\.Web\.UI\.WebControls\.SqlDataSource | System\.Web\.UI\.WebControls\.SqlDataSourceView | Microsoft\.Whos\.Framework\.Data\.SqlUtility`),
		},
	}
}

func NewCsharpAndWeakRandomNumberGenerator() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "4b546b8d-0d0c-4b37-ad5f-f8f788019a3e",
			Name:        "Weak Random Number Generator",
			Description: "The use of a predictable random value can lead to vulnerabilities when used in certain security critical contexts. For more information access: (https://security-code-scan.github.io/#SCS0005).",
			Severity:    severity.Low.ToString(),
			Confidence:  confidence.High.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`new Random\(\)`),
			regexp.MustCompile(`new byte\[.*\]`),
			regexp.MustCompile(`GetBytes\(\)`),
		},
	}
}

func NewCsharpAndWeakCipherOrCBCOrECBMode() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "f9436bdd-8a80-4168-98fa-17af9f8c51b8",
			Name:        "Weak Cipher Mode",
			Description: "The cipher provides no way to detect that the data has been tampered with. If the cipher text can be controlled by an attacker, it could be altered without detection. The use of AES in CBC mode with a HMAC is recommended guaranteeing integrity and confidentiality. For more information access: (https://security-code-scan.github.io/#SCS0013).",
			Severity:    severity.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`(using)(([^O]|O[^r]|Or[^g]|Org[^.]|Org\.[^B]|Org\.B[^o]|Org\.Bo[^u]|Org\.Bou[^n]|Org\.Boun[^c]|Org\.Bounc[^y]|Org\.Bouncy[^C]|Org\.BouncyC[^a]|Org\.BouncyCa[^s]|Org\.BouncyCas[^t]|Org\.BouncyCast[^l]|Org\.BouncyCastl[^e])*)(\);)`),
			regexp.MustCompile(`CreateEncryptor\(.*\)`),
			regexp.MustCompile(`new CryptoStream\(.*\)`),
			regexp.MustCompile(`Write\(.*\)`),
			regexp.MustCompile(`new BinaryWriter\(.*\)`),
		},
	}
}

func NewCsharpAndFormsAuthenticationCookielessMode() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "81f4b46e-164c-478f-ad3a-df263d73c00c",
			Name:        "Forms Authentication Cookieless Mode",
			Description: "Authentication cookies should not be sent in the URL. Doing so allows attackers to gain unauthorized access to authentication tokens (web server logs, referrer headers, and browser history) and more easily perform session fixation / hijacking attacks. For more information checkout the CWE-598 (https://cwe.mitre.org/data/definitions/598.html) advisory.",
			Severity:    severity.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`\<authentication\s*mode\s*=\s*["|']Forms`),
			regexp.MustCompile(`(\<forms)((([^c]|c[^o]|co[^o]|coo[^k]|cook[^i]|cooki[^e]|cookie[^l]|cookiel[^e]|cookiele[^s]|cookieles[^s])*)|([^U]|U[^s]|Us[^e]|Use[^C]|UseC[^o]|UseCo[^o]|UseCoo[^k]|UseCook[^i]|UseCooki[^e]|UseCookie[^s])*)(\/\>)`),
		},
	}
}

func NewCsharpAndFormsAuthenticationCrossAppRedirects() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "05cd1279-54a7-4770-9527-3ecd6c83b167",
			Name:        "Forms Authentication Cross App Redirects",
			Description: "Enabling cross-application redirects can allow unvalidated redirect attacks via the returnUrl parameter during the login process. Disable cross-application redirects to by setting the enableCrossAppRedirects attribute to false. For more information checkout the CWE-601 (https://cwe.mitre.org/data/definitions/601.html) advisory.",
			Severity:    severity.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`\<authentication\s*mode\s*=\s*["|']Forms`),
			regexp.MustCompile(`\<forms`),
			regexp.MustCompile(`enableCrossAppRedirects\s*=\s*["|']true`),
		},
	}
}

func NewCsharpAndFormsAuthenticationWeakCookieProtection() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "97df49ee-1043-4158-ba33-efed0ac3a28d",
			Name:        "Forms Authentication Weak Cookie Protection",
			Description: "Forms Authentication cookies must use strong encryption and message authentication code (MAC) validation to protect the cookie value from inspection and tampering. Configure the forms element’s protection attribute to All to enable cookie data validation and encryption. For more information checkout the CWE-565 (https://cwe.mitre.org/data/definitions/565.html) advisory.",
			Severity:    severity.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`\<authentication\s*mode\s*=\s*["|']Forms`),
			regexp.MustCompile(`\<forms`),
			regexp.MustCompile(`protection\s*=\s*["|'](None|Encryption|Validation)`),
		},
	}
}

func NewCsharpAndFormsAuthenticationWeakTimeout() text.TextRule {
	return text.TextRule{
		Metadata: engine.Metadata{
			ID:          "fc00e79f-03ae-43eb-967f-46d8b6efa2ea",
			Name:        "Forms Authentication Weak Timeout",
			Description: "Excessive authentication timeout values provide attackers with a large window of opportunity to hijack user’s authentication tokens. For more information checkout the CWE-613 (https://cwe.mitre.org/data/definitions/613.html) advisory.",
			Severity:    severity.Medium.ToString(),
			Confidence:  confidence.Low.ToString(),
		},
		Type: text.AndMatch,
		Expressions: []*regexp.Regexp{
			regexp.MustCompile(`\<authentication\s*mode\s*=\s*["|']Forms`),
			regexp.MustCompile(`\<forms`),
			regexp.MustCompile(`timeout\s*=\s*["|'](1[6-9]|[2-9][0-9]*)`),
		},
	}
}
