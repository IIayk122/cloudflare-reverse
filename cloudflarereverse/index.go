package cloudflarereverse

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/0xF7A4C6/clean-http/cleanhttp"
	fp "github.com/0xF7A4C6/fingerprint-client/fingerprintclient"
	http "github.com/bogdanfinn/fhttp"
	"io/ioutil"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var (
	re = regexp.MustCompile(`[0-9]*\.[0-9]+:[0-9]+:`)
)

// Need improvement
func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func randHexString(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func encryptFp(password string) string {
	input := `{"0":["TEMPORARY","innerHeight","innerWidth","length","pageXOffset","pageYOffset","screenLeft","screenTop","screenX","screenY","scrollX","scrollY","n.maxTouchPoints"],"1":["PERSISTENT","d.DOCUMENT_POSITION_DISCONNECTED","d.ELEMENT_NODE","d.childElementCount"],"2":["d.ATTRIBUTE_NODE","d.DOCUMENT_POSITION_PRECEDING"],"3":["d.TEXT_NODE"],"4":["d.CDATA_SECTION_NODE","d.DOCUMENT_POSITION_FOLLOWING"],"5":["d.ENTITY_REFERENCE_NODE"],"6":["d.ENTITY_NODE"],"7":["d.PROCESSING_INSTRUCTION_NODE"],"8":["n.deviceMemory","d.COMMENT_NODE","d.DOCUMENT_POSITION_CONTAINS"],"9":["d.DOCUMENT_NODE","d.nodeType"],"10":["d.DOCUMENT_TYPE_NODE"],"11":["d.DOCUMENT_FRAGMENT_NODE"],"12":["d.NOTATION_NODE"],"16":["n.hardwareConcurrency","d.DOCUMENT_POSITION_CONTAINED_BY"],"32":["d.DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC"],"1040":["outerHeight"],"1920":["outerWidth"],"N":["AbortController","AbortSignal","AbsoluteOrientationSensor","AbstractRange","Accelerometer","AggregateError","AnalyserNode","Animation","AnimationEffect","AnimationEvent","AnimationPlaybackEvent","AnimationTimeline","ArrayBuffer","Attr","Audio","AudioBuffer","AudioBufferSourceNode","AudioContext","AudioData","AudioDecoder","AudioDestinationNode","AudioEncoder","AudioListener","AudioNode","AudioParam","AudioParamMap","AudioProcessingEvent","AudioScheduledSourceNode","AudioWorklet","AudioWorkletNode","AuthenticatorAssertionResponse","AuthenticatorAttestationResponse","AuthenticatorResponse","BackgroundFetchManager","BackgroundFetchRecord","BackgroundFetchRegistration","BarProp","BaseAudioContext","BatteryManager","BeforeInstallPromptEvent","BeforeUnloadEvent","BigInt","BigInt64Array","BigUint64Array","BiquadFilterNode","Blob","BlobEvent","Bluetooth","BluetoothCharacteristicProperties","BluetoothDevice","BluetoothRemoteGATTCharacteristic","BluetoothRemoteGATTDescriptor","BluetoothRemoteGATTServer","BluetoothRemoteGATTService","BluetoothUUID","Boolean","BroadcastChannel","BrowserCaptureMediaStreamTrack","ByteLengthQueuingStrategy","CDATASection","CSSAnimation","CSSConditionRule","CSSContainerRule","CSSCounterStyleRule","CSSFontFaceRule","CSSFontPaletteValuesRule","CSSGroupingRule","CSSImageValue","CSSImportRule","CSSKeyframeRule","CSSKeyframesRule","CSSKeywordValue","CSSLayerBlockRule","CSSLayerStatementRule","CSSMathClamp","CSSMathInvert","CSSMathMax","CSSMathMin","CSSMathNegate","CSSMathProduct","CSSMathSum","CSSMathValue","CSSMatrixComponent","CSSMediaRule","CSSNamespaceRule","CSSNumericArray","CSSNumericValue","CSSPageRule","CSSPerspective","CSSPositionValue","CSSPropertyRule","CSSRotate","CSSRule","CSSRuleList","CSSScale","CSSSkew","CSSSkewX","CSSSkewY","CSSStyleDeclaration","CSSStyleRule","CSSStyleSheet","CSSStyleValue","CSSSupportsRule","CSSTransformComponent","CSSTransformValue","CSSTransition","CSSTranslate","CSSUnitValue","CSSUnparsedValue","CSSVariableReferenceValue","Cache","CacheStorage","CanvasCaptureMediaStreamTrack","CanvasFilter","CanvasGradient","CanvasPattern","CanvasRenderingContext2D","ChannelMergerNode","ChannelSplitterNode","CharacterData","Clipboard","ClipboardEvent","ClipboardItem","CloseEvent","Comment","CompositionEvent","CompressionStream","ConstantSourceNode","ConvolverNode","CookieChangeEvent","CookieStore","CookieStoreManager","CountQueuingStrategy","Credential","CredentialsContainer","CropTarget","Crypto","CryptoKey","CustomElementRegistry","CustomEvent","CustomStateSet","DOMError","DOMException","DOMImplementation","DOMMatrix","DOMMatrixReadOnly","DOMParser","DOMPoint","DOMPointReadOnly","DOMQuad","DOMRect","DOMRectList","DOMRectReadOnly","DOMStringList","DOMStringMap","DOMTokenList","DataTransfer","DataTransferItem","DataTransferItemList","DataView","Date","DecompressionStream","DelayNode","DelegatedInkTrailPresenter","DeviceMotionEvent","DeviceMotionEventAcceleration","DeviceMotionEventRotationRate","DeviceOrientationEvent","Document","DocumentFragment","DocumentTimeline","DocumentType","DragEvent","DynamicsCompressorNode","Element","ElementInternals","EncodedAudioChunk","EncodedVideoChunk","Error","ErrorEvent","EvalError","Event","EventCounts","EventSource","EventTarget","External","EyeDropper","FeaturePolicy","FederatedCredential","File","FileList","FileReader","FileSystemDirectoryHandle","FileSystemFileHandle","FileSystemHandle","FileSystemWritableFileStream","FinalizationRegistry","Float32Array","Float64Array","FocusEvent","FontData","FontFace","FontFaceSetLoadEvent","FormData","FormDataEvent","FragmentDirective","Function","GainNode","Gamepad","GamepadButton","GamepadEvent","GamepadHapticActuator","Geolocation","GeolocationCoordinates","GeolocationPosition","GeolocationPositionError","GravitySensor","Gyroscope","HID","HIDConnectionEvent","HIDDevice","HIDInputReportEvent","HTMLAllCollection","HTMLAnchorElement","HTMLAreaElement","HTMLAudioElement","HTMLBRElement","HTMLBaseElement","HTMLBodyElement","HTMLButtonElement","HTMLCanvasElement","HTMLCollection","HTMLDListElement","HTMLDataElement","HTMLDataListElement","HTMLDetailsElement","HTMLDialogElement","HTMLDirectoryElement","HTMLDivElement","HTMLDocument","HTMLElement","HTMLEmbedElement","HTMLFieldSetElement","HTMLFontElement","HTMLFormControlsCollection","HTMLFormElement","HTMLFrameElement","HTMLFrameSetElement","HTMLHRElement","HTMLHeadElement","HTMLHeadingElement","HTMLHtmlElement","HTMLIFrameElement","HTMLImageElement","HTMLInputElement","HTMLLIElement","HTMLLabelElement","HTMLLegendElement","HTMLLinkElement","HTMLMapElement","HTMLMarqueeElement","HTMLMediaElement","HTMLMenuElement","HTMLMetaElement","HTMLMeterElement","HTMLModElement","HTMLOListElement","HTMLObjectElement","HTMLOptGroupElement","HTMLOptionElement","HTMLOptionsCollection","HTMLOutputElement","HTMLParagraphElement","HTMLParamElement","HTMLPictureElement","HTMLPreElement","HTMLProgressElement","HTMLQuoteElement","HTMLScriptElement","HTMLSelectElement","HTMLSlotElement","HTMLSourceElement","HTMLSpanElement","HTMLStyleElement","HTMLTableCaptionElement","HTMLTableCellElement","HTMLTableColElement","HTMLTableElement","HTMLTableRowElement","HTMLTableSectionElement","HTMLTemplateElement","HTMLTextAreaElement","HTMLTimeElement","HTMLTitleElement","HTMLTrackElement","HTMLUListElement","HTMLUnknownElement","HTMLVideoElement","HashChangeEvent","Headers","Highlight","HighlightRegistry","History","IDBCursor","IDBCursorWithValue","IDBDatabase","IDBFactory","IDBIndex","IDBKeyRange","IDBObjectStore","IDBOpenDBRequest","IDBRequest","IDBTransaction","IDBVersionChangeEvent","IIRFilterNode","IdleDeadline","IdleDetector","Image","ImageBitmap","ImageBitmapRenderingContext","ImageCapture","ImageData","ImageDecoder","ImageTrack","ImageTrackList","Ink","InputDeviceCapabilities","InputDeviceInfo","InputEvent","Int16Array","Int32Array","Int8Array","IntersectionObserver","IntersectionObserverEntry","Keyboard","KeyboardEvent","KeyboardLayoutMap","KeyframeEffect","LargestContentfulPaint","LaunchParams","LaunchQueue","LayoutShift","LayoutShiftAttribution","LinearAccelerationSensor","Location","Lock","LockManager","MIDIAccess","MIDIConnectionEvent","MIDIInput","MIDIInputMap","MIDIMessageEvent","MIDIOutput","MIDIOutputMap","MIDIPort","Map","MediaCapabilities","MediaDeviceInfo","MediaDevices","MediaElementAudioSourceNode","MediaEncryptedEvent","MediaError","MediaKeyMessageEvent","MediaKeySession","MediaKeyStatusMap","MediaKeySystemAccess","MediaKeys","MediaList","MediaMetadata","MediaQueryList","MediaQueryListEvent","MediaRecorder","MediaSession","MediaSource","MediaStream","MediaStreamAudioDestinationNode","MediaStreamAudioSourceNode","MediaStreamEvent","MediaStreamTrack","MediaStreamTrackEvent","MediaStreamTrackGenerator","MediaStreamTrackProcessor","MessageChannel","MessageEvent","MessagePort","MimeType","MimeTypeArray","MouseEvent","MutationEvent","MutationObserver","MutationRecord","NamedNodeMap","NavigateEvent","Navigation","NavigationCurrentEntryChangeEvent","NavigationDestination","NavigationHistoryEntry","NavigationPreloadManager","NavigationTransition","Navigator","NavigatorManagedData","NavigatorUAData","NetworkInformation","Node","NodeFilter","NodeIterator","NodeList","Notification","Number","OTPCredential","Object","OfflineAudioCompletionEvent","OfflineAudioContext","OffscreenCanvas","OffscreenCanvasRenderingContext2D","Option","OrientationSensor","OscillatorNode","OverconstrainedError","PageTransitionEvent","PannerNode","PasswordCredential","Path2D","PaymentAddress","PaymentInstruments","PaymentManager","PaymentMethodChangeEvent","PaymentRequest","PaymentRequestUpdateEvent","PaymentResponse","Performance","PerformanceElementTiming","PerformanceEntry","PerformanceEventTiming","PerformanceLongTaskTiming","PerformanceMark","PerformanceMeasure","PerformanceNavigation","PerformanceNavigationTiming","PerformanceObserver","PerformanceObserverEntryList","PerformancePaintTiming","PerformanceResourceTiming","PerformanceServerTiming","PerformanceTiming","PeriodicSyncManager","PeriodicWave","PermissionStatus","Permissions","PictureInPictureEvent","PictureInPictureWindow","Plugin","PluginArray","PointerEvent","PopStateEvent","Presentation","PresentationAvailability","PresentationConnection","PresentationConnectionAvailableEvent","PresentationConnectionCloseEvent","PresentationConnectionList","PresentationReceiver","PresentationRequest","ProcessingInstruction","Profiler","ProgressEvent","Promise","PromiseRejectionEvent","Proxy","PublicKeyCredential","PushManager","PushSubscription","PushSubscriptionOptions","RTCCertificate","RTCDTMFSender","RTCDTMFToneChangeEvent","RTCDataChannel","RTCDataChannelEvent","RTCDtlsTransport","RTCEncodedAudioFrame","RTCEncodedVideoFrame","RTCError","RTCErrorEvent","RTCIceCandidate","RTCIceTransport","RTCPeerConnection","RTCPeerConnectionIceErrorEvent","RTCPeerConnectionIceEvent","RTCRtpReceiver","RTCRtpSender","RTCRtpTransceiver","RTCSctpTransport","RTCSessionDescription","RTCStatsReport","RTCTrackEvent","RadioNodeList","Range","RangeError","ReadableByteStreamController","ReadableStream","ReadableStreamBYOBReader","ReadableStreamBYOBRequest","ReadableStreamDefaultController","ReadableStreamDefaultReader","ReferenceError","RegExp","RelativeOrientationSensor","RemotePlayback","ReportingObserver","Request","ResizeObserver","ResizeObserverEntry","ResizeObserverSize","Response","SVGAElement","SVGAngle","SVGAnimateElement","SVGAnimateMotionElement","SVGAnimateTransformElement","SVGAnimatedAngle","SVGAnimatedBoolean","SVGAnimatedEnumeration","SVGAnimatedInteger","SVGAnimatedLength","SVGAnimatedLengthList","SVGAnimatedNumber","SVGAnimatedNumberList","SVGAnimatedPreserveAspectRatio","SVGAnimatedRect","SVGAnimatedString","SVGAnimatedTransformList","SVGAnimationElement","SVGCircleElement","SVGClipPathElement","SVGComponentTransferFunctionElement","SVGDefsElement","SVGDescElement","SVGElement","SVGEllipseElement","SVGFEBlendElement","SVGFEColorMatrixElement","SVGFEComponentTransferElement","SVGFECompositeElement","SVGFEConvolveMatrixElement","SVGFEDiffuseLightingElement","SVGFEDisplacementMapElement","SVGFEDistantLightElement","SVGFEDropShadowElement","SVGFEFloodElement","SVGFEFuncAElement","SVGFEFuncBElement","SVGFEFuncGElement","SVGFEFuncRElement","SVGFEGaussianBlurElement","SVGFEImageElement","SVGFEMergeElement","SVGFEMergeNodeElement","SVGFEMorphologyElement","SVGFEOffsetElement","SVGFEPointLightElement","SVGFESpecularLightingElement","SVGFESpotLightElement","SVGFETileElement","SVGFETurbulenceElement","SVGFilterElement","SVGForeignObjectElement","SVGGElement","SVGGeometryElement","SVGGradientElement","SVGGraphicsElement","SVGImageElement","SVGLength","SVGLengthList","SVGLineElement","SVGLinearGradientElement","SVGMPathElement","SVGMarkerElement","SVGMaskElement","SVGMatrix","SVGMetadataElement","SVGNumber","SVGNumberList","SVGPathElement","SVGPatternElement","SVGPoint","SVGPointList","SVGPolygonElement","SVGPolylineElement","SVGPreserveAspectRatio","SVGRadialGradientElement","SVGRect","SVGRectElement","SVGSVGElement","SVGScriptElement","SVGSetElement","SVGStopElement","SVGStringList","SVGStyleElement","SVGSwitchElement","SVGSymbolElement","SVGTSpanElement","SVGTextContentElement","SVGTextElement","SVGTextPathElement","SVGTextPositioningElement","SVGTitleElement","SVGTransform","SVGTransformList","SVGUnitTypes","SVGUseElement","SVGViewElement","Sanitizer","Scheduler","Scheduling","Screen","ScreenDetailed","ScreenDetails","ScreenOrientation","ScriptProcessorNode","SecurityPolicyViolationEvent","Selection","Sensor","SensorErrorEvent","Serial","SerialPort","ServiceWorker","ServiceWorkerContainer","ServiceWorkerRegistration","Set","ShadowRoot","SharedWorker","SourceBuffer","SourceBufferList","SpeechSynthesisErrorEvent","SpeechSynthesisEvent","SpeechSynthesisUtterance","StaticRange","StereoPannerNode","Storage","StorageEvent","StorageManager","String","StylePropertyMap","StylePropertyMapReadOnly","StyleSheet","StyleSheetList","SubmitEvent","SubtleCrypto","Symbol","SyncManager","SyntaxError","TaskAttributionTiming","TaskController","TaskPriorityChangeEvent","TaskSignal","Text","TextDecoder","TextDecoderStream","TextEncoder","TextEncoderStream","TextEvent","TextMetrics","TextTrack","TextTrackCue","TextTrackCueList","TextTrackList","TimeRanges","Touch","TouchEvent","TouchList","TrackEvent","TransformStream","TransformStreamDefaultController","TransitionEvent","TreeWalker","TrustedHTML","TrustedScript","TrustedScriptURL","TrustedTypePolicy","TrustedTypePolicyFactory","TypeError","UIEvent","URIError","URL","URLPattern","URLSearchParams","USB","USBAlternateInterface","USBConfiguration","USBConnectionEvent","USBDevice","USBEndpoint","USBInTransferResult","USBInterface","USBIsochronousInTransferPacket","USBIsochronousInTransferResult","USBIsochronousOutTransferPacket","USBIsochronousOutTransferResult","USBOutTransferResult","Uint16Array","Uint32Array","Uint8Array","Uint8ClampedArray","UserActivation","VTTCue","ValidityState","VideoColorSpace","VideoDecoder","VideoEncoder","VideoFrame","VideoPlaybackQuality","VirtualKeyboard","VirtualKeyboardGeometryChangeEvent","VisualViewport","WakeLock","WakeLockSentinel","WaveShaperNode","WeakMap","WeakRef","WeakSet","WebGL2RenderingContext","WebGLActiveInfo","WebGLBuffer","WebGLContextEvent","WebGLFramebuffer","WebGLProgram","WebGLQuery","WebGLRenderbuffer","WebGLRenderingContext","WebGLSampler","WebGLShader","WebGLShaderPrecisionFormat","WebGLSync","WebGLTexture","WebGLTransformFeedback","WebGLUniformLocation","WebGLVertexArrayObject","WebKitCSSMatrix","WebKitMutationObserver","WebSocket","WebTransport","WebTransportBidirectionalStream","WebTransportDatagramDuplexStream","WebTransportError","WheelEvent","Window","WindowControlsOverlay","WindowControlsOverlayGeometryChangeEvent","Worker","Worklet","WritableStream","WritableStreamDefaultController","WritableStreamDefaultWriter","XMLDocument","XMLHttpRequest","XMLHttpRequestEventTarget","XMLHttpRequestUpload","XMLSerializer","XPathEvaluator","XPathExpression","XPathResult","XRAnchor","XRAnchorSet","XRBoundedReferenceSpace","XRCPUDepthInformation","XRDOMOverlayState","XRDepthInformation","XRFrame","XRHitTestResult","XRHitTestSource","XRInputSource","XRInputSourceArray","XRInputSourceEvent","XRInputSourcesChangeEvent","XRLayer","XRLightEstimate","XRLightProbe","XRPose","XRRay","XRReferenceSpace","XRReferenceSpaceEvent","XRRenderState","XRRigidTransform","XRSession","XRSessionEvent","XRSpace","XRSystem","XRTransientInputHitTestResult","XRTransientInputHitTestSource","XRView","XRViewerPose","XRViewport","XRWebGLBinding","XRWebGLDepthInformation","XRWebGLLayer","XSLTProcessor","addEventListener","alert","atob","blur","btoa","cancelAnimationFrame","cancelIdleCallback","captureEvents","clearInterval","clearTimeout","close","confirm","createImageBitmap","decodeURI","decodeURIComponent","dispatchEvent","encodeURI","encodeURIComponent","escape","eval","fetch","find","focus","getComputedStyle","getScreenDetails","getSelection","isFinite","isNaN","matchMedia","moveBy","moveTo","open","openDatabase","parseFloat","parseInt","postMessage","print","prompt","queryLocalFonts","queueMicrotask","releaseEvents","removeEventListener","reportError","requestAnimationFrame","requestIdleCallback","resizeBy","resizeTo","scroll","scrollBy","scrollTo","setInterval","setTimeout","showDirectoryPicker","showOpenFilePicker","showSaveFilePicker","stop","structuredClone","unescape","webkitCancelAnimationFrame","webkitMediaStream","webkitRTCPeerConnection","webkitRequestAnimationFrame","webkitRequestFileSystem","webkitResolveLocalFileSystemURL","webkitSpeechGrammar","webkitSpeechGrammarList","webkitSpeechRecognition","webkitSpeechRecognitionError","webkitSpeechRecognitionEvent","webkitURL","n.canShare","n.clearAppBadge","n.getBattery","n.getGamepads","n.getInstalledRelatedApps","n.getUserMedia","n.javaEnabled","n.registerProtocolHandler","n.requestMIDIAccess","n.requestMediaKeySystemAccess","n.sendBeacon","n.setAppBadge","n.share","n.unregisterProtocolHandler","n.vibrate","n.webkitGetUserMedia","d.addEventListener","d.adoptNode","d.append","d.appendChild","d.captureEvents","d.caretRangeFromPoint","d.clear","d.cloneNode","d.close","d.compareDocumentPosition","d.contains","d.createAttribute","d.createAttributeNS","d.createCDATASection","d.createComment","d.createDocumentFragment","d.createElement","d.createElementNS","d.createEvent","d.createExpression","d.createNSResolver","d.createNodeIterator","d.createProcessingInstruction","d.createRange","d.createTextNode","d.createTreeWalker","d.dispatchEvent","d.elementFromPoint","d.elementsFromPoint","d.evaluate","d.execCommand","d.exitFullscreen","d.exitPictureInPicture","d.exitPointerLock","d.getAnimations","d.getElementById","d.getElementsByClassName","d.getElementsByName","d.getElementsByTagName","d.getElementsByTagNameNS","d.getRootNode","d.getSelection","d.hasChildNodes","d.hasFocus","d.importNode","d.insertBefore","d.isDefaultNamespace","d.isEqualNode","d.isSameNode","d.lookupNamespaceURI","d.lookupPrefix","d.normalize","d.open","d.prepend","d.queryCommandEnabled","d.queryCommandIndeterm","d.queryCommandState","d.queryCommandSupported","d.queryCommandValue","d.querySelector","d.querySelectorAll","d.releaseEvents","d.removeChild","d.removeEventListener","d.replaceChild","d.replaceChildren","d.webkitCancelFullScreen","d.webkitExitFullscreen","d.write","d.writeln"],"C":["Array"],"o":["Atomics","CSS","Intl","JSON","Math","Reflect","WebAssembly","caches","chrome","clientInformation","console","cookieStore","crypto","customElements","document","external","frameElement","frames","globalThis","history","indexedDB","launchQueue","localStorage","location","locationbar","menubar","navigation","navigator","parent","performance","personalbar","scheduler","screen","scrollbars","self","sessionStorage","speechSynthesis","statusbar","styleMedia","toolbar","top","trustedTypes","visualViewport","webkitStorageInfo","window","n.bluetooth","n.clipboard","n.connection","n.credentials","n.geolocation","n.hid","n.ink","n.keyboard","n.locks","n.managed","n.mediaCapabilities","n.mediaDevices","n.mediaSession","n.mimeTypes","n.permissions","n.plugins","n.presentation","n.scheduling","n.serial","n.serviceWorker","n.storage","n.usb","n.userActivation","n.userAgentData","n.virtualKeyboard","n.wakeLock","n.webkitPersistentStorage","n.webkitTemporaryStorage","n.windowControlsOverlay","n.xr","d.activeElement","d.anchors","d.applets","d.body","d.childNodes","d.children","d.defaultView","d.documentElement","d.embeds","d.featurePolicy","d.firstChild","d.firstElementChild","d.fonts","d.forms","d.fragmentDirective","d.head","d.images","d.implementation","d.lastChild","d.lastElementChild","d.links","d.location","d.plugins","d.scripts","d.scrollingElement","d.styleSheets","d.timeline"],"Infinity":["Infinity"],"NaN":["NaN"],"false":["closed","crossOriginIsolated","originAgentCluster","n.webdriver","d.fullscreen","d.hidden","d.prerendering","d.wasDiscarded","d.webkitHidden","d.webkitIsFullScreen","d.xmlStandalone"],"0.8999999761581421":["devicePixelRatio"],"u":["event","undefined"],"true":["isSecureContext","offscreenBuffering","n.cookieEnabled","n.onLine","n.pdfViewerEnabled","d.fullscreenEnabled","d.isConnected","d.pictureInPictureEnabled","d.webkitFullscreenEnabled"],"x":["onabort","onafterprint","onanimationend","onanimationiteration","onanimationstart","onappinstalled","onauxclick","onbeforeinput","onbeforeinstallprompt","onbeforematch","onbeforeprint","onbeforeunload","onbeforexrselect","onblur","oncancel","oncanplay","oncanplaythrough","onchange","onclick","onclose","oncontextlost","oncontextmenu","oncontextrestored","oncuechange","ondblclick","ondevicemotion","ondeviceorientation","ondeviceorientationabsolute","ondrag","ondragend","ondragenter","ondragleave","ondragover","ondragstart","ondrop","ondurationchange","onemptied","onended","onerror","onfocus","onformdata","ongotpointercapture","onhashchange","oninput","oninvalid","onkeydown","onkeypress","onkeyup","onlanguagechange","onload","onloadeddata","onloadedmetadata","onloadstart","onlostpointercapture","onmessage","onmessageerror","onmousedown","onmouseenter","onmouseleave","onmousemove","onmouseout","onmouseover","onmouseup","onmousewheel","onoffline","ononline","onpagehide","onpageshow","onpause","onplay","onplaying","onpointercancel","onpointerdown","onpointerenter","onpointerleave","onpointermove","onpointerout","onpointerover","onpointerrawupdate","onpointerup","onpopstate","onprogress","onratechange","onrejectionhandled","onreset","onresize","onscroll","onsearch","onsecuritypolicyviolation","onseeked","onseeking","onselect","onselectionchange","onselectstart","onslotchange","onstalled","onstorage","onsubmit","onsuspend","ontimeupdate","ontoggle","ontransitioncancel","ontransitionend","ontransitionrun","ontransitionstart","onunhandledrejection","onunload","onvolumechange","onwaiting","onwebkitanimationend","onwebkitanimationiteration","onwebkitanimationstart","onwebkittransitionend","onwheel","opener","n.doNotTrack","d.all","d.currentScript","d.doctype","d.fullscreenElement","d.nextSibling","d.nodeValue","d.onabort","d.onanimationend","d.onanimationiteration","d.onanimationstart","d.onauxclick","d.onbeforecopy","d.onbeforecut","d.onbeforeinput","d.onbeforematch","d.onbeforepaste","d.onbeforexrselect","d.onblur","d.oncancel","d.oncanplay","d.oncanplaythrough","d.onchange","d.onclick","d.onclose","d.oncontextlost","d.oncontextmenu","d.oncontextrestored","d.oncopy","d.oncuechange","d.oncut","d.ondblclick","d.ondrag","d.ondragend","d.ondragenter","d.ondragleave","d.ondragover","d.ondragstart","d.ondrop","d.ondurationchange","d.onemptied","d.onended","d.onerror","d.onfocus","d.onformdata","d.onfreeze","d.onfullscreenchange","d.onfullscreenerror","d.ongotpointercapture","d.oninput","d.oninvalid","d.onkeydown","d.onkeypress","d.onkeyup","d.onload","d.onloadeddata","d.onloadedmetadata","d.onloadstart","d.onlostpointercapture","d.onmousedown","d.onmouseenter","d.onmouseleave","d.onmousemove","d.onmouseout","d.onmouseover","d.onmouseup","d.onmousewheel","d.onpaste","d.onpause","d.onplay","d.onplaying","d.onpointercancel","d.onpointerdown","d.onpointerenter","d.onpointerleave","d.onpointerlockchange","d.onpointerlockerror","d.onpointermove","d.onpointerout","d.onpointerover","d.onpointerrawupdate","d.onpointerup","d.onprerenderingchange","d.onprogress","d.onratechange","d.onreadystatechange","d.onreset","d.onresize","d.onresume","d.onscroll","d.onsearch","d.onsecuritypolicyviolation","d.onseeked","d.onseeking","d.onselect","d.onselectionchange","d.onselectstart","d.onslotchange","d.onstalled","d.onsubmit","d.onsuspend","d.ontimeupdate","d.ontoggle","d.ontransitioncancel","d.ontransitionend","d.ontransitionrun","d.ontransitionstart","d.onvisibilitychange","d.onvolumechange","d.onwaiting","d.onwebkitanimationend","d.onwebkitanimationiteration","d.onwebkitanimationstart","d.onwebkitfullscreenchange","d.onwebkitfullscreenerror","d.onwebkittransitionend","d.onwheel","d.ownerDocument","d.parentElement","d.parentNode","d.pictureInPictureElement","d.pointerLockElement","d.previousSibling","d.rootElement","d.textContent","d.webkitCurrentFullScreenElement","d.webkitFullscreenElement","d.xmlEncoding","d.xmlVersion"],"https://discord.com":["origin"],"Mozilla":["n.appCodeName"],"Netscape":["n.appName"],"5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36":["n.appVersion"],"fr-FR":["n.language"],"fr-FR,fr,en-US,en":["n.languages"],"Win32":["n.platform"],"Gecko":["n.product"],"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36":["n.userAgent"],"Google Inc.":["n.vendor"],"about:blank":["d.URL","d.documentURI","d.referrer"],"":["d.adoptedStyleSheets"],"https://discord.com/":["d.baseURI"],"UTF-8":["d.characterSet","d.charset","d.inputEncoding"],"BackCompat":["d.compatMode"],"text/html":["d.contentType"],"s":["d.cookie"],"off":["d.designMode"],"discord.com":["d.domain"],"10/22/2022 01:40:47":["d.lastModified"],"#document":["d.nodeName"],"complete":["d.readyState"],"visible":["d.visibilityState","d.webkitVisibilityState"]}`

	compressed := strings.Split(Compress(input, password), "===")[0]

	return compressed
}

func GetCfbm(brFp *fp.Fingerprint, proxy string) (string, error) {
	header := http.Header{
		`authority`:          {`discord.com`},
		`accept`:             {`*/*`},
		`accept-language`:    {`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		`content-type`:       {`application/json`},
		`origin`:             {`https://discord.com`},
		`sec-ch-ua`:          {`"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`},
		`sec-ch-ua-mobile`:   {`?0`},
		`sec-ch-ua-platform`: {fmt.Sprintf(`"%s"`, brFp.Navigator.Platform)},
		`sec-fetch-dest`:     {`empty`},
		`sec-fetch-mode`:     {`cors`},
		`sec-fetch-site`:     {`same-origin`},
		`user-agent`:         {brFp.Navigator.UserAgent},

		http.HeaderOrderKey: {
			"authority",
			"accept",
			"accept-language",
			"content-type",
			"origin",
			"sec-ch-ua",
			"sec-ch-ua-mobile",
			"sec-ch-ua-platform",
			"sec-fetch-dest",
			"sec-fetch-mode",
			"sec-fetch-site",
			"user-agent",
		},
	}

	client, err := cleanhttp.NewCleanHttpClient(&cleanhttp.Config{
		Proxy:     proxy,
		BrowserFp: brFp,
	})
	if err != nil {
		return "", err
	}

	resp, err := client.Do(cleanhttp.RequestOption{
		Method: "GET",
		Url:    "https://discord.com",
		Header: header,
	})
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	Cf := &CfParams{
		R: strings.Split(strings.Split(string(response), "r:'")[1], "',m")[0],
		M: strings.Split(strings.Split(string(response), "m:'")[1], "'};")[0],
	}

	resp, err = client.Do(cleanhttp.RequestOption{
		Method: "GET",
		Url:    "https://discord.com/cdn-cgi/challenge-platform/h/g/scripts/jsd/5da7637f/invisible.js",
		Header: header,
	})
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	JsScript := string(response)

	var Pass string
	for _, char := range strings.Split(strings.Split(JsScript, "'.split(';')")[0], ";") {
		if len(char) == 65 {
			Pass = char
		}
	}

	if Pass == "" {
		return "", fmt.Errorf("cant find encryption password")
	}

	Base := re.FindString(JsScript)
	S := Base + strings.Split(JsScript, Base)[1][:43]

	timing := float64(randInt(100, 400))

	jsonPayload, _ := json.Marshal(FingerprintPayload{
		Src: "worker",
		T:   float64(timing+float64(randInt(300, 500))) + rand.Float64(),
		S:   S,
		Fp: Fingerprint{
			Results: []string{
				randHexString(16),
			},
			Timing: int(timing),
		},
		M:  Cf.M,
		Wp: encryptFp(Pass),
	})

	resp, err = client.Do(cleanhttp.RequestOption{
		Method: "POST",
		Url:    fmt.Sprintf("https://discord.com/cdn-cgi/challenge-platform/h/b/cv/result/%s", Cf.R),
		Header: header,
		Body:   strings.NewReader(string(jsonPayload)),
	})
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("cant submit answer")
	}

	for _, c := range client.CookieJar.Cookies(&url.URL{Host: "discord.com"}) {
		if c.Name == "__cf_bm" {
			return c.Value, nil
		}
	}

	return "", fmt.Errorf("no cookie found")
}