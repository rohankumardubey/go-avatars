package bottts

import (
	"sort"
)

var eyes = map[string]string{
	"bulging": `
	<path fill-rule="evenodd" clip-rule="evenodd" d="M28 42C37.9411 42 46 33.9411 46 24C46 14.0589 37.9411 6 28 6C18.0589 6 10 14.0589 10 24C10 33.9411 18.0589 42 28 42Z" fill="black" fill-opacity="0.2"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M74 42C83.9411 42 92 33.9411 92 24C92 14.0589 83.9411 6 74 6C64.0589 6 56 14.0589 56 24C56 33.9411 64.0589 42 74 42Z" fill="black" fill-opacity="0.2"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M28 39C36.2843 39 43 32.2843 43 24C43 15.7157 36.2843 9 28 9C19.7157 9 13 15.7157 13 24C13 32.2843 19.7157 39 28 39Z" fill="#F1EEDA"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M74 39C82.2843 39 89 32.2843 89 24C89 15.7157 82.2843 9 74 9C65.7157 9 59 15.7157 59 24C59 32.2843 65.7157 39 74 39Z" fill="#F1EEDA"/>
	<rect x="26" y="15" width="10" height="10" rx="2" fill="black" fill-opacity="0.8"/>
	<rect x="74" y="15" width="10" height="10" rx="2" fill="black" fill-opacity="0.8"/>
`,
	"dizzy": `
	<path fill-rule="evenodd" clip-rule="evenodd" d="M25 27.2L30.5 32.7C31 33.1 31.7 33.1 32.1 32.7L33.7 31.1C34.1 30.6 34.1 29.9 33.7 29.5L28.2 24L33.7 18.5C34.1 18 34.1 17.3 33.7 16.9L32.1 15.3C31.6 14.9 30.9 14.9 30.5 15.3L25 20.8L19.5 15.3C19 14.9 18.3 14.9 17.9 15.3L16.3 16.9C15.9 17.3 15.9 18 16.3 18.5L21.8 24L16.3 29.5C15.9 30 15.9 30.7 16.3 31.1L17.9 32.7C18.4 33.1 19.1 33.1 19.5 32.7L25 27.2Z" fill="black" fill-opacity="0.8"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M79 27.2L84.5 32.7C85 33.1 85.7 33.1 86.1 32.7L87.7 31.1C88.1 30.6 88.1 29.9 87.7 29.5L82.2 24L87.7 18.5C88.1 18 88.1 17.3 87.7 16.9L86.1 15.3C85.6 14.9 84.9 14.9 84.5 15.3L79 20.8L73.5 15.3C73 14.9 72.3 14.9 71.9 15.3L70.3 16.9C69.9 17.3 69.9 18 70.3 18.5L75.8 24L70.3 29.5C69.9 30 69.9 30.7 70.3 31.1L71.9 32.7C72.4 33.1 73.1 33.1 73.5 32.7L79 27.2Z" fill="black" fill-opacity="0.8"/>
`,
	"eva": `
	<path fill-rule="evenodd" clip-rule="evenodd" d="M53 0C87.7469 0 102.001 17.4742 102 31C101.999 44.5258 82.4108 48 53 48C23.9528 48 2 44.5258 2 31C2 17.4742 17.1142 0 53 0Z" fill="black" fill-opacity="0.8"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M28.8179 34.654C22.2912 33.3001 17.5833 28.3121 18.3026 23.513C19.0218 18.7139 24.8959 15.9211 31.4226 17.275C37.9493 18.629 42.6572 23.617 41.9379 28.416C41.2187 33.2151 35.3446 36.0079 28.8179 34.654Z" fill="#25A6F5"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M75.4226 34.654C68.8959 36.0079 63.0218 33.2151 62.3026 28.416C61.5833 23.617 66.2912 18.629 72.8179 17.275C79.3446 15.9211 85.2187 18.7139 85.9379 23.513C86.6572 28.3121 81.9493 33.3001 75.4226 34.654Z" fill="#25A6F5"/>
`,
	"frame-1": `
	<rect y="4" width="104" height="42" rx="10" fill="black" fill-opacity="0.8"/>
	<rect x="28" y="25" width="10" height="11" rx="2" fill="#8BDDFF"/>
	<rect x="66" y="25" width="10" height="11" rx="2" fill="#8BDDFF"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M21 4H29L12 46H4L21 4Z" fill="white" fill-opacity="0.4"/>
`,
	"frame-2": `
	<rect x="8" y="10" width="88" height="36" rx="4" fill="black" fill-opacity="0.8"/>
	<rect x="28" y="21" width="10" height="17" rx="2" fill="#5EF2B8"/>
	<rect x="66" y="21" width="10" height="17" rx="2" fill="#5EF2B8"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M83 10H88L76 46H71L83 10Z" fill="white" fill-opacity="0.4"/>
`,
	"glow": `
	<path fill-rule="evenodd" clip-rule="evenodd" d="M21 45C29.2843 45 36 38.2843 36 30C36 21.7157 29.2843 15 21 15C12.7157 15 6 21.7157 6 30C6 38.2843 12.7157 45 21 45Z" fill="white" fill-opacity="0.1"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M83 45C91.2843 45 98 38.2843 98 30C98 21.7157 91.2843 15 83 15C74.7157 15 68 21.7157 68 30C68 38.2843 74.7157 45 83 45Z" fill="white" fill-opacity="0.1"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M21 42C27.6274 42 33 36.6274 33 30C33 23.3726 27.6274 18 21 18C14.3726 18 9 23.3726 9 30C9 36.6274 14.3726 42 21 42Z" fill="white" fill-opacity="0.1"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M83 42C89.6274 42 95 36.6274 95 30C95 23.3726 89.6274 18 83 18C76.3726 18 71 23.3726 71 30C71 36.6274 76.3726 42 83 42Z" fill="white" fill-opacity="0.1"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M21 36C24.3137 36 27 33.3137 27 30C27 26.6863 24.3137 24 21 24C17.6863 24 15 26.6863 15 30C15 33.3137 17.6863 36 21 36Z" fill="white" fill-opacity="0.8"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M83 36C86.3137 36 89 33.3137 89 30C89 26.6863 86.3137 24 83 24C79.6863 24 77 26.6863 77 30C77 33.3137 79.6863 36 83 36Z" fill="white" fill-opacity="0.8"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M21 33C22.6569 33 24 31.6569 24 30C24 28.3431 22.6569 27 21 27C19.3431 27 18 28.3431 18 30C18 31.6569 19.3431 33 21 33Z" fill="white" fill-opacity="0.8"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M83 33C84.6569 33 86 31.6569 86 30C86 28.3431 84.6569 27 83 27C81.3431 27 80 28.3431 80 30C80 31.6569 81.3431 33 83 33Z" fill="white" fill-opacity="0.8"/>
`,
	"hal": `
	<path fill-rule="evenodd" clip-rule="evenodd" d="M52 48C65.2548 48 76 37.2548 76 24C76 10.7452 65.2548 0 52 0C38.7452 0 28 10.7452 28 24C28 37.2548 38.7452 48 52 48Z" fill="white" fill-opacity="0.4"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M52 44C63.0457 44 72 35.0457 72 24C72 12.9543 63.0457 4 52 4C40.9543 4 32 12.9543 32 24C32 35.0457 40.9543 44 52 44Z" fill="black" fill-opacity="0.8"/>
	<path d="M64.5726 15.8153C64.8743 16.2779 65.4939 16.4082 65.9565 16.1064C66.419 15.8046 66.5494 15.185 66.2476 14.7225L64.5726 15.8153ZM61.5815 9.95547C61.1256 9.64384 60.5033 9.76084 60.1917 10.2168C59.88 10.6728 59.997 11.295 60.453 11.6067L61.5815 9.95547ZM56.3568 9.64222C56.8854 9.80237 57.4437 9.50373 57.6039 8.97518C57.764 8.44663 57.4654 7.88832 56.9368 7.72816L56.3568 9.64222ZM45.7206 8.19769C45.2074 8.40179 44.9569 8.98326 45.161 9.49645C45.3651 10.0096 45.9465 10.2602 46.4597 10.0561L45.7206 8.19769ZM41.7603 13.0388C42.1638 12.6617 42.1852 12.0289 41.8081 11.6254C41.431 11.2219 40.7981 11.2005 40.3947 11.5776L41.7603 13.0388ZM36.4567 17.1052C36.2325 17.6099 36.4599 18.2008 36.9646 18.425C37.4694 18.6492 38.0603 18.4218 38.2845 17.9171L36.4567 17.1052ZM66.2476 14.7225C65.0212 12.8427 63.433 11.2208 61.5815 9.95547L60.453 11.6067C62.0875 12.7238 63.49 14.1559 64.5726 15.8153L66.2476 14.7225ZM56.9368 7.72816C55.3733 7.25438 53.7155 7 52.0001 7V9C53.5169 9 54.9793 9.2248 56.3568 9.64222L56.9368 7.72816ZM52.0001 7C49.784 7 47.6646 7.42456 45.7206 8.19769L46.4597 10.0561C48.1724 9.37496 50.0413 9 52.0001 9V7ZM40.3947 11.5776C38.7378 13.1261 37.3906 15.0029 36.4567 17.1052L38.2845 17.9171C39.108 16.0633 40.2968 14.4066 41.7603 13.0388L40.3947 11.5776Z" fill="white" fill-opacity="0.9"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M52 34C57.5228 34 62 29.5228 62 24C62 18.4772 57.5228 14 52 14C46.4772 14 42 18.4772 42 24C42 29.5228 46.4772 34 52 34Z" fill="#C6080C"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M52 28C54.2091 28 56 26.2091 56 24C56 21.7909 54.2091 20 52 20C49.7909 20 48 21.7909 48 24C48 26.2091 49.7909 28 52 28Z" fill="#EE9337"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M52 25C52.5523 25 53 24.5523 53 24C53 23.4477 52.5523 23 52 23C51.4477 23 51 23.4477 51 24C51 24.5523 51.4477 25 52 25Z" fill="#F5F94F"/>
`,
	"happy": `
	<path d="M18 19L30 17" stroke="black" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
	<path d="M20 31C20 27.686 22.9098 25 27 25C30.0902 25 33 27.686 33 31" stroke="black" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
	<path d="M86 20L74 17" stroke="black" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
	<path d="M84 31C84 27.686 81.0902 25 78 25C73.9098 25 71 27.686 71 31" stroke="black" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
`,
	"hearts": `
	<path d="M29.2691 9.67983C26.7216 9.81334 24.305 11.9225 23.0195 13.8332C21.5357 12.0676 18.9175 10.2223 16.3701 10.3558C10.8883 10.6431 7.51531 14.1586 7.74073 18.4598C8.0406 24.1816 12.6244 27.3464 17.4365 30.7046C19.1523 31.8531 22.441 34.8494 22.8627 35.5214C23.2845 36.1935 25.0034 36.1278 25.4425 35.3862C25.8817 34.6446 28.7463 31.3502 30.3355 30.0286C34.7674 26.1859 38.9981 22.5592 38.6982 16.8374C38.4728 12.5362 34.7509 9.39254 29.2691 9.67983Z" fill="#FF5353" fill-opacity="0.8"/>
	<path d="M87.6299 10.3558C85.0825 10.2223 82.4586 12.0673 80.9805 13.8332C79.6893 11.9222 77.2784 9.81331 74.7309 9.67981C69.2491 9.39252 65.5272 12.5361 65.3017 16.8374C65.0019 22.5591 69.2297 26.1857 73.6645 30.0286C75.2508 31.3501 78.2083 34.6738 78.5575 35.3862C78.9067 36.0987 80.623 36.2131 81.1373 35.5214C81.6515 34.8298 84.8449 31.8529 86.5635 30.7046C91.3728 27.3462 95.9594 24.1816 96.2593 18.4598C96.4847 14.1586 93.1117 10.6431 87.6299 10.3558Z" fill="#FF5353" fill-opacity="0.8"/>
`,
	"round-frame-1": `
	<rect y="12" width="104" height="32" rx="16" fill="black" fill-opacity="0.8"/>
	<rect x="24" y="22" width="10" height="12" rx="2" fill="#F4F4F4"/>
	<rect x="70" y="22" width="10" height="12" rx="2" fill="#F4F4F4"/>
`,
	"round-frame-2": `
	<rect y="11" width="104" height="34" rx="17" fill="black" fill-opacity="0.8"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M29 41C36.1797 41 42 35.1797 42 28C42 20.8203 36.1797 15 29 15C21.8203 15 16 20.8203 16 28C16 35.1797 21.8203 41 29 41Z" fill="#F1EEDA"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M75 41C82.1797 41 88 35.1797 88 28C88 20.8203 82.1797 15 75 15C67.8203 15 62 20.8203 62 28C62 35.1797 67.8203 41 75 41Z" fill="#F1EEDA"/>
	<rect x="24" y="23" width="10" height="10" rx="2" fill="black" fill-opacity="0.8"/>
	<rect x="70" y="23" width="10" height="10" rx="2" fill="black" fill-opacity="0.8"/>
`,
}

var eyeNames []string

func init() {
	eyeNames = make([]string, 0, len(eyes))
	for k := range eyes {
		eyeNames = append(eyeNames, k)
	}
	sort.Strings(eyeNames)
}
