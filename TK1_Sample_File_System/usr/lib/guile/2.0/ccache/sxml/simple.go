GOOF----LE-4-2.0,S      ] × 4       hI      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  sxml¤	g  simple¤	 ¤		g  filenameS¤	
f  sxml/simple.scm¤	g  importsS¤	g  ssax¤	g  input-parse¤	 ¤	 ¤	 ¤	 ¤	g  	transform¤	 ¤	 ¤	g  ice-9¤	g  match¤	 ¤	 ¤	g  srfi¤	g  srfi-13¤	 ¤	 ¤	 ¤	g  exportsS¤	g  	xml->sxml¤	 g  	sxml->xml¤	!g  sxml->string¤	" ! ¤	#g  set-current-module¤	$# ¤	%# ¤	&g  string-concatenate/shared¤	'g  string?¤	(g  ssax:reverse-collect-str¤	)g  
next-token¤	*] ¤	+f  reading internal DOCTYPE¤	,g  peek-next-char¤	-g  	read-char¤	.f  ]¤	/g  read-internal-doctype-as-string¤	0g  
namespacesS¤	10¤	2g  declare-namespaces?S¤	32	¤	4g  trim-whitespace?S¤	54	¤	6g  entitiesS¤	76	¤	8g  default-entity-handlerS¤	98	¤	:g  doctype-handlerS¤	;:	¤	<13579; ¤	=g  current-input-port¤	>g  map¤	?g  ssax:uri-string->symbol¤	@g  error¤	A@ ¤	B@ ¤	Cf  no matching pattern¤	Dg  ssax:scan-Misc¤	ED ¤	FD ¤	Gg  eof-object?¤	HG ¤	IG ¤	Jg  parser-error¤	KJ ¤	LJ ¤	Mf  XML [22], unexpected EOF¤	Ng  PI¤	Og  *PI*¤	Pg  ssax:read-pi-body-as-string¤	Qg  DECL¤	Rg  string->symbol¤	SR ¤	TR ¤	Uf  DOCTYPE¤	Vf  .XML [22], expected DOCTYPE declaration, found ¤	Wg  assert-curr-char¤	XW ¤	YW ¤	Zg  ssax:S-chars¤	[Z ¤	\Z ¤	]f  XML [28], space after DOCTYPE¤	^g  ssax:skip-S¤	_^ ¤	`^ ¤	ag  ssax:read-QName¤	ba ¤	ca ¤	dg  ssax:ncname-starting-char?¤	ed ¤	fd ¤	gg  ssax:read-external-id¤	hg ¤	ig ¤	j>[ ¤	kf  XML [28], end-of-DOCTYPE¤	lg  call-with-values¤	mg  ssax:skip-internal-dtd¤	n6
¤	on1 ¤	pg  append¤	qg  assq¤	rg  	*DEFAULT*¤	sg  START¤	tf  XML [22], unexpected markup ¤	ug  ssax:Prefix-XML¤	vu ¤	wu ¤	xf  space¤	yg  ssax:complete-start-tag¤	zy ¤	{y ¤	|g  	EMPTY-TAG¤	}g   ssax:reverse-collect-str-drop-ws¤	~g  attlist-fold¤	g  symbol-append¤ f  :¤ g  @¤ g  EMPTY¤ g  ssax:assert-token¤  ¤  ¤ g  ssax:read-markup-token¤  ¤  ¤ g  END¤ f  [elementvalid] broken for ¤ f   while expecting ¤ g  assoc¤  ¤  ¤ f  preserve¤ g  ssax:read-char-data¤  ¤  ¤ g  string-null?¤ f  [GIMatch] broken for ¤ g  
ENTITY-REF¤ g  ssax:handle-parsed-entity¤  ¤  ¤ g  PCDATA¤ f  * with char content only; unexpected token ¤ f  XML [43] broken for ¤ g  open-input-string¤ g  reverse¤ g  *TOP*¤ g  make-hash-table¤  g  	hashq-ref¤ ¡g  symbol->string¤ ¢g  string-index¤ £g  	substring¤ ¤f  "Invalid QName: more than one colon¤ ¥g  for-each¤ ¦g  char-alphabetic?¤ §g  
string-ref¤ ¨f  Invalid name starting character¤ ©g  string-for-each¤ ªf  0123456789.-_¤ «f  Invalid name character¤ ¬g  
hashq-set!¤ ­g  
check-name¤ ®g  attribute-value->xml¤ ¯g  string->escaped-xml¤ °g  
procedure?¤ ±g  with-output-to-port¤ ²g  call-with-output-string¤ ³g  display¤ ´f  ="¤ µg  attribute->xml¤ ¶f  bad attribute¤ ·f  bad attributes¤ ¸f  </¤ ¹f  >¤ ºf  bad element body¤ »f   />¤ ¼g  element->xml¤ ½g  entity->xml¤ ¾f  <?¤ ¿f  ?>¤ Àg  pi->xml¤ Ág  current-output-port¤ Âg  *ENTITY*¤ Ãf  bad *ENTITY* args¤ Äg  length¤ Åf  bad *PI* args¤ Æg  string-concatenate-reverse¤ Çg  foldts¤ Èg  append!¤ Ég  list->char-set¤ Êg  car¤ Ëg  string-length¤ Ìg  assv¤ Íg  make-char-quotator¤ Îf  &lt;¤ Ï<Î¤ Ðf  &gt;¤ Ñ>Ð¤ Òf  &amp;¤ Ó&Ò¤ Ôf  &quot;¤ Õ"Ô¤ ÖÏÑÓÕ ¤C 5  hÐA    ]4	
"5 4% >  "  G   &'    h   ¥  ] (  C (   C"  c(  (  C45C45$  "ÿÿÅ(  "  
45"ÿÿ "ÿÿ             g  	fragments
	  g  	fragments		{ g  result			{ g  strs			{  g  filenamef  sxml/simple.scm
	,
		-				.			/			-			1			2		$	4		(	5		1	5		3	6		8	6		:	6		>	2		A	7		F	7	/	I	7	)	S	7		V	9		Y	;		_	<		f	=		o	=		p	:		q	>		{	9		{	1		~	1	/		1	: 	1	 	   g  nameg  ssax:reverse-collect-str C(R&)*+,-.     hH   Ú   ]4L5 >4L5$  4L>  "  G    C 4L 5 C      Ò       g  fragment
		B  g  filenamef  sxml/simple.scm
	B			D			D			D		
	D	$		D			C			E			E			E			G		5	H		:	I		;	I	 	@	I	 		B
  g  nameg  loop C   h    ¡   ]	O  Q 45 6         g  port
		 g  loop		  g  filenamef  sxml/simple.scm
	@
		B			A	 		  g  nameg  read-internal-doctype-as-string C/R<=>?BC        h8      ] $  #  L $  "  45C6           g  el
		4 g  w		, g  x			,  g  filenamef  sxml/simple.scm
	l				m			o		#	q		*	o		0	m	 		4   C     h      ]L O  6 x       g  
namespaces
		  g  filenamef  sxml/simple.scm
	k			l	 		  g  nameg  munge-namespaces CFILMNOPQTUVY\]`cfijkl/m     h@   C   ] L$  LL LL$  4L5"  6L$  4L>  "  G  "   D ;       g  filenamef  sxml/simple.scm
 	 		@
   Copqr hp     -  /     0   3  #   #  4 L$  4L5$  L"  LL"  L544L54LL55L D           g  entities
		k g  
namespaces		k  g  filenamef  sxml/simple.scm
 		 		# 	/	' 		1	x		2	y		6	y		:	y		>	x		F	z		I	z		T 		U 		X 		_	u		g 		k 	 		k

g  entitiesS
g  
namespacesS    Cs     h   C   ] L $  L 6D        ;       g  filenamef  sxml/simple.scm
 	 		
   Copqr hp     -  /     0   3  #   #  4 L$  4L5$  L"  LL"  L544L54LL55L D           g  entities
		k g  
namespaces		k  g  filenamef  sxml/simple.scm
 		 		# 	/	' 		1	x		2	y		6	y		:	y		>	x		F	z		I	z		T 		U 		X 		_	u		g 		k 	 		k

g  entitiesS
g  
namespacesS    CtFILMNOPst 
h   +  ]4 545$   6$  34 5   "ÿÿ£$  L 6 	6#      g  port
	  g  elems	  g  entities		  g  
namespaces		  g  seed		  g  token			  g  key		!  g  target		/	E g  seed		E	_  	g  filenamef  sxml/simple.scm
 	 	 	  g  nameg  #scan-for-significant-prolog-token-2 CwTx{|}(~R hH   ×   ] $  $4455"    C       Ï       g  attr
		A g  accum		A g  name			7 g  w			. g  x			.  g  filenamef  sxml/simple.scm
 		
	~		 		 		" 	-	$ 		( 		: 	 			A	   CR`L     h      ]L  6      {       g  token
		 g  exp-kind		 g  exp-head			  g  filenamef  sxml/simple.scm
 	 			   CR   hH   ×   ] $  $4455"    C       Ï       g  attr
		A g  accum		A g  name			7 g  w			. g  x			.  g  filenamef  sxml/simple.scm
 		
	~		 		 		" 	-	$ 		( 		: 	 			A	   C    h       ]45$   C C    x       g  string1
		 g  string2		 g  seed			  g  filenamef  sxml/simple.scm
 	 			   CIL        h      ]LL 6      {       g  token
		 g  exp-kind		 g  exp-head			  g  filenamef  sxml/simple.scm
 	 			   C}(~R     hH   ×   ] $  $4455"    C       Ï       g  attr
		A g  accum		A g  name			7 g  w			. g  x			.  g  filenamef  sxml/simple.scm
 		
	~		 		 		" 	-	$ 		( 		: 	 			A	   CRNOP     h   ~   ]L  6   v       g  port
		 g  entities		 g  seed			  g  filenamef  sxml/simple.scm
 	 			   C    h       ]45$   C C    x       g  string1
		 g  string2		 g  seed			  g  filenamef  sxml/simple.scm
 	 			   CsL h  Ê  ]<4 >  G 45$  C$  4L O >  "  G  L$  45"  454	
L5L	$  $L	L	
	4	45
5	
"  L	(  "  LC$  /4 5   "ÿÿ	$  -4 L O 5  "ÿþÓ$  NL&  4 L	>  "  G  "   4L LL5  "ÿþ| 6Â      g  port
	 g  entities	 g  expect-eof?		 g  seed		 g  seed		 g  
term-token		 g  key		+ g  seed		w Á g  attrs		w Á g  w		  ¢ g  x	
  ¢ g  target	 Í ã g  seed	 ã ù g  seed	/ g  seed	p  g  filenamef  sxml/simple.scm
 		~	~	  	  	  	-  	  	 ® 	 			  g  nameg  loop C   h  ³  ]F4 L>  G 		$  jL $  
45"  4545
$  $44	55"  (  
"  

C	$  4<45$  45"   O >  "  G  L $  
45"  4545
$  $44	55"  (  
"  

C4M5

$  

"  

O 

L	 L Q 6 «      g  start-tag-head
	 g  port	 g  entities		 g  
namespaces		 g  preserve-ws?		 g  parent-seed		 g  elem-gi		 g  
attributes		 g  
namespaces		 g  expected-content			 g  seed	
	I  g  attrs		I  g  w		Z	t g  x		Z	t g  seed	
 ñ; g  attrs	 ñ; g  w	 g  x	 g  t	
D[ g  preserve-ws?	
[ g  loop	n  g  filenamef  sxml/simple.scm
 		P	~		_ 		d 		h 	-	j 		n 	  	 ø	~	 	 	 	- 	 	( 	 		  g  nameg  handle-start-tag C  h8   2  ]H45KO L Q  6 *      g  start-tag-head
		7 g  port		7 g  elems			7 g  entities			7 g  
namespaces			7 g  preserve-ws?			7 g  seed			7 g  xml-space-gi			7 g  handle-start-tag			7  	g  filenamef  sxml/simple.scm
 	 		7	  g  nameg  element-parser C   hÈ  o  ]r" 4545$  6$  '45 "ÿÿ¯$  å4	
5		$  "  4>  "  G  	4>  "  G  4>  "  G  45	4455$  45"  
4>  "  G  [454	
LO LLLLO >  G 6$  @4LO LLLLO >  G 
		
66O O Q L Q  "ÿþE       g      g  port
	Á g  seed	Á g  port		 g  seed		 g  token		 g  key		% g  target		3	I g  seed		I	W g  
token-head		cE g  t			o  g  docname		 ÃE g  systemid	
 áE g  internal-subset?	E g  elems	,E g  entities	,E g  
namespaces	,E g  seed	,E g  elems	q g  entities		q g  
namespaces	
q g  seed	q g  #scan-for-significant-prolog-token-2	©Á g  element-parser	©Á  g  filenamef  sxml/simple.scm
 	 	Á	  g  nameg  parser C' 	     h°   =  -  /    0   3  #  45  #  #  #  #  #  #  O O Q Q 4 5$  4 5"   	44	55

C  5      g  string-or-port
	 ® g  
namespaces	 ® g  declare-namespaces?		 ® g  trim-whitespace?		 ® g  entities		 ® g  default-entity-handler		 ® g  doctype-handler		 ® g  munge-namespaces		k ® g  parser		k ® g  port		  ® g  elements	
 ¦ ®  g  filenamef  sxml/simple.scm
	V
		V	/	'	W	 	B	Z	  Ó	  Ó	  Ô	  Ó	  Ö	  Ö	 ¢ Ö	) ¤ Ö	 ¦ Ö	 ¦ Ó	 ª ×	 	 ®
g  
namespacesSg  declare-namespaces?S	g  trim-whitespace?S	g  entitiesS	g  default-entity-handlerS	g  doctype-handlerS	   g  nameg  	xml->sxmlg  documentationf  Use SSAX to parse an XML document into SXML. Takes one optional
argument, @var{string-or-port}, which defaults to the current input
port. CR4i5   ¡¢£@¤¥¦§@¨©¦¢ª@«       h8   ·   ]	4 5$  C4 5$  C LL 6    ¯       g  c
		4 g  t			4 g  t		4  g  filenamef  sxml/simple.scm
 ê		 ë			 ë		 ë	0	 ë	>	 ë	0	 ë		, ì	"	4 ì	 
		4   C      hh   å   ] $  X44 
55$  "  &4 
5_$  "  4 L 5$  L  O  6CC      Ý       g  s
		b g  t		K g  t		-	H  g  filenamef  sxml/simple.scm
 ä		 å			 æ		 æ	*	 æ		 æ		# ç		- ç		- æ		; è		? è		E è		O å		^ é	 		b   C¬ 
     hÈ   »  ]!4L  5$  C4 54:5$  4
5"  $  "  $  45"  $  1445:5$  4 >  "  G  "   "   4 O  >  "  G  	L  6    ³      g  name
	 Ä g  str	 Ä g  i		# Ä g  t		:	M g  head		M Ä g  tail		d Ä  g  filenamef  sxml/simple.scm
 Û		 Ü		 Ü		 Ý		 Ý	
	 Þ		# Ý	
	+ ß		, ß	"	: ß		M Ý	
	U à		V à		] à	-	_ à		d Ý	
	l á		m á		p á	!	w á	0	y á	!	} á	  á	  â	  â	  â	  ã	 ® î	 ³ ã	 Ä ï	 	 Ä   C O  ­R®'¯°±²³       h   _   ]L  6      W       g  port
		
  g  filenamef  sxml/simple.scm
		
	- 		
   C      hh     ] $  4 >  "  G   6 (  C4 5$   64 5$   64 O 56       g  value
		g g  port		g  g  filenamef  sxml/simple.scm
 õ
	 ÷			 ö		
 ø		 ø		 ø		# ù		' ù		- ö		0 ü		: ö		B ý		C þ		M ö		U ÿ		X		g	 		g	  g  nameg  attribute-value->xml C®R­³´®        hX   Ô   ]4 >  "  G  4 >  "  G  4>  "  G  4>  "  G  "6Ì       g  attr
		X g  value		X g  port			X  g  filenamef  sxml/simple.scm

					)		-		4		=			X
	 			X	  g  nameg  attribute->xml CµR­³µ@¶· ¸¹º»       hp  ¾  ]4 >  "  G  4<>  "  G  4 >  "  G  $  "  $  ^4 >  "  G  $  4>  "  G  "  4 >  "  G  "ÿÿ(  "  "4 >  "  G  "  "ÿÿo"   $  4>>  "  G  "  e$  4>  "  G  "ÿÿÛ(  04>  "  G  4 >  "  G  	6
 6"ÿÿ6 ¶      g  tag
	o g  attrs	o g  body		o g  port		o g  attrs		F Ï g  attr		P « g  body	 ú_  g  filenamef  sxml/simple.scm

					)		B		F		I		M		P		P		S		i		m		n		s	"	v	-	}	 	 	 	 ¢	 ¨	 ±	 ¶	 º	 Ã	 Ï	 Þ	 â	 ã	 ú	 ý 		
!	!	!	"	"	%	
&$	*$	1$	:%	Q&	U&	Y(	_(	_	k)	o)	 3	o	  g  nameg  element->xml C¼R³ h8      ]4&>  "  G  4 >  "  G  ;6             g  name
		2 g  port		2  g  filenamef  sxml/simple.scm
,
	-		.		2/	 		2	  g  nameg  entity->xml C½R³¾¿ h`   Ó   ]4>  "  G  4 >  "  G  4 >  "  G  4>  "  G  6      Ë       g  tag
		Z g  str		Z g  port			Z  g  filenamef  sxml/simple.scm
2
	3		3		3		4		+5		?6		V7		Z7	 
		Z	  g  nameg  pi->xml CÀRÁ Â@ÃÄ½OÅÀ¼¥        h   \   ] L 6      T       g  x
		
  g  filenamef  sxml/simple.scm
T		
T	 		
   C'¯°±²³    h   _   ]L  6      W       g  port
		
  g  filenamef  sxml/simple.scm
^		
^	- 		
   C      hp  ë  - . , 3 #  45  $  ú $  å $  	 6$  5"  	 6 $  4 5$  
 6"ÿÿÓ"ÿÿÏ	$  ;"  	
 6 $  "4 5	$    6"ÿÿÍ"ÿÿÉ $  $$  &  "  "  "  $  "  6O  64 5$   6 (  C $  . &  C4 5$   64 O 56C     ã      g  tree
	k g  port	k g  tag		+ g  elems	 Ä g  attrs	 ò  g  filenamef  sxml/simple.scm
9
	9	*	>		 =		#?		$?		(?		+A		+A		6B	
	;D		?D		HB	
	PH		SH	,	UH		UF		XF		YF		]F		^F	,	cF	4	eF	,	gF	)	kF		pG		uG	 B	
 L	 L	( L	 J	 J	 J	 J	 J	, ¡J	4 £J	, ¦J	) ªJ	 ¯K	 ³K	& ¹K	 ÄN	 ÄN	 ÉO	  ÍO	 ÐO	5 ÑO	. ÕO	 ×P	% ÚP	( ßO	 âQ	  òN	 R	'R	1R	T	U	%=	-V	3=	EZ	O=	W[	Z^	i]	 E	k  g  nameg  	sxml->xmlg  documentationf  Serialize the sxml tree @var{tree} as XML. The output will be written
to the current output port, unless the optional argument @var{port} is
present. C RÆÇ   h   q   ]C    i       g  seed
		 g  tree		  g  filenamef  sxml/simple.scm
f		g	 			   CÈ h      ] 6             g  seed
		
 g  kid-seed		
 g  tree			
  g  filenamef  sxml/simple.scm
h		
i	 		
	   C'   h      ]45$   C C  {       g  seed
		 g  tree		  g  filenamef  sxml/simple.scm
j		k	
	k		k	 			   C  h   õ   ]4 56    í       g  sxml
		  g  filenamef  sxml/simple.scm
a
	e		l		e		d	 		  g  nameg  sxml->stringg  documentationf  MDetag an sxml tree @var{sxml} into a string. Does not perform any
formatting. C!RÉ>Ê¢ËÌ§³£ hÐ   ð  ]*4 L 
5$  ®"  4 5$  C$  s44 5L54 L 5$  !44 5>  "  G  "   4>  "  G  "ÿÿw4 4 556
"ÿÿV 6      è      g  str
	 Ê g  port	 Ê g  bad-pos		 Ê g  from		 · g  to		 · g  i		B	R g  quoted-char		R ¡ g  new-to		R ¡  g  filenamef  sxml/simple.scm
y		v		z		{		}				"		&~		/		2	"	>		?		B	'	B		Ev		R		[		_		`		c		t	 	 	 ¡	 ¤	 «	- ³	 ·	 ·}	 Ê|	 	 Ê	   C     h    ·   ]	44 55 O C    ¯       g  char-encoding
		 g  	bad-chars		  g  filenamef  sxml/simple.scm
o
	p		p	#	p		p	 		  g  nameg  make-char-quotator CÍR4ÍiÖ5¯RC
      g  m
		, g  *good-cache*
%G*k  g  filenamef  sxml/simple.scm		
x	,
	@
%@	V
%A Ú	%G Ú	*n Ù
, õ
-ß
29
3,
4f2
:þ9
=øa
AÂo
AÃ	AÉ	AË	AÎ
 	AÐ
   C6 