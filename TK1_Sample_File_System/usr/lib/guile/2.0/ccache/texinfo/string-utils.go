GOOF----LE-4-2.0¯A      ] W 4       hñ      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  texinfo¤	g  string-utils¤	 ¤		g  filenameS¤	
f  texinfo/string-utils.scm¤	g  importsS¤	g  srfi¤	g  srfi-13¤	 ¤	 ¤	g  srfi-14¤	 ¤	 ¤	g  oop¤	g  goops¤	 ¤	 ¤	 ¤	g  exportsS¤	g  escape-special-chars¤	g  transform-string¤	g  expand-tabs¤	g  center-string¤	g  left-justify-string¤	g  right-justify-string¤	g  collapse-repeated-chars¤	 g  make-text-wrapper¤	!g  fill-string¤	"g  string->wrapped-lines¤	# !" 
¤	$g  set-current-module¤	%$ ¤	&$ ¤	'g  open-output-string¤	(g  char?¤	)g  char=?¤	*g  
procedure?¤	+g  string?¤	,g  string-index¤	-g  boolean?¤	.g  throw¤	/g  bad-type¤	0f  'expected #t, char, string, or procedure¤	1g  display¤	2g  string-length¤	3g  	substring¤	4g  string-for-each¤	5g  
write-char¤	6g  get-output-string¤	7g  make-string¤	8g  string¤	9g  max¤	:g  string-append¤	;g  odd?¤	<f   ¤	=g  char-set-complement¤	>g  char-set:whitespace¤	?g  reverse¤	@g  split-by-single-words¤	Ag  
string-ref¤	Bg  end-of-sentence?¤	Cg  
line-widthS¤	DC
¤	Eg  expand-tabs?S¤	FE¤	Gg  	tab-widthS¤	HG	¤	Ig  collapse-whitespace?S¤	JI	¤	Kg  subsequent-indentS¤	LK	¤	Mg  initial-indentS¤	NM	¤	Og  break-long-words?S¤	PO	¤	QDFHJLNP ¤	Rg  string-trim¤	Sf   ¤	Tg  string-join¤	Uf  
¤	Vg  infix¤C 5  hx:  Å   ]4	
#5 4& >  "  G   '() h   c   ]L  6      [       g  c
		
  g  filenamef  texinfo/string-utils.scm
	P		
	P	% 		
   C*+,   h   c   ]L  6      [       g  c
		
  g  filenamef  texinfo/string-utils.scm
	T		
	T	% 		
   C- h   [   ]L C   S       g  c
		  g  filenamef  texinfo/string-utils.scm
	V	 		   C./01 h   k   ]4L 5L 6 c       g  c
		  g  filenamef  texinfo/string-utils.scm
	Y			Y	,		Y	# 		   C1       h   c   ]LL 6      [       g  c
		
  g  filenamef  texinfo/string-utils.scm
	Z		
	Z	# 		
   C21345       h    {   ]4L 5$  L 6 L 6     s       g  c
		  g  filenamef  texinfo/string-utils.scm
	b			c			c			e			h	 		   C6        hX  -	  - . , 3 #  #  45 45$  O "  F45$  "  545$  O "  45$  O "  	4	
545$  O "  	O $  24 5$   44 
5>  "  G  "   "   4O  $  "  
$  "  4 5>  "  G  $  14 5$  44 5>  "  G  "   "   6%	      g  str
	X g  match?	X g  replace		X g  start		X g  end		X g  os		#X g  matcher	 X g  replacer	 ¤X  g  filenamef  texinfo/string-utils.scm
	(
		N		#	N		&	O		0	O		<	Q		F	O		M	S		W	O		c	U		m	O		y	W		}	W	%		W	/ 	W	 	N	 	X	 	X	 ¤	N	 ¬	]	 ¯	]	 ¶	]	 º	]	 »	^	 ¾	^	 Î	^	 ß	a	 ô	j		k		k		a	!	n	$	n	+	n	/	n	0	o	3	o	B	o	X	q	 (	X		  g  nameg  transform-stringg  documentationf ÀUses @var{match?} against each character in @var{str}, and performs a
replacement on each character for which matches are found.

@var{match?} may either be a function, a character, a string, or
@code{#t}.  If @var{match?}  is a function, then it takes a single
character as input, and should return @samp{#t} for matches.
@var{match?} is a character, it is compared to each string character
using @code{char=?}.  If @var{match?} is a string, then any character
in that string will be considered a match.  @code{#t} will cause 
every character to be a match.

If @var{replace} is a function, it is called with the matched
character as an argument, and the returned value is sent to the output
string via @samp{display}.  If @var{replace} is anything else, it is
sent through the output string via @samp{display}.

Note that te replacement for the matched characters does not need to
be a single character.  That is what differentiates this function from
@samp{string-map}, and what makes it useful for applications such as
converting @samp{#\&} to @samp{"&amp;"} in web page text.  Some other
functions in this module are just wrappers around common uses of
@samp{transform-string}.  Transformations not possible with this
function should probably be done with regular expressions.

If @var{start} and @var{end} are given, they control which portion
of the string undergoes transformation.  The entire input string
is still output, though.  So, if @var{start} is @samp{5}, then the
first five characters of @var{str} will still appear in the returned
string.

@lisp
; these two are equivalent...
 (transform-string str #\space #\-) ; change all spaces to -'s
 (transform-string str (lambda (c) (char=? #\space c)) #\-)
@end lisp CR7        h(   u  - . , 3 #  	 	4 56 m      g  str
		' g  tab-size		'  g  filenamef  texinfo/string-utils.scm
	s
		{		'	y	 		'  g  nameg  expand-tabsg  documentationf  ÅReturns a copy of @var{str} with all tabs expanded to spaces.  @var{tab-size} defaults to 8.

Assuming tab size of 8, this is equivalent to: @lisp
 (transform-string str #\tab "        ")
@end lisp CR()     h   e   ] L 6      ]       g  c
		
  g  filenamef  texinfo/string-utils.scm
 		
 	$ 		
   C,     h   e   ]L  6      ]       g  c
		
  g  filenamef  texinfo/string-utils.scm
 		
 	$ 		
   C8     h   e   ]L  6      ]       g  c
		
  g  filenamef  texinfo/string-utils.scm
 		
 	  		
   C        h0   )  ] 45$  O "  O O 6    !      g  str
		, g  special-chars		, g  escape-char			,  g  filenamef  texinfo/string-utils.scm
	}
	 		 		, 	 		,	  g  nameg  escape-special-charsg  documentationf GReturns a copy of @var{str} with all given special characters preceded
by the given @var{escape-char}.

@var{special-chars} can either be a single character, or a string consisting
of all the special characters.

@lisp
;; make a string regexp-safe...
 (escape-special-chars "***(Example String)***"  
                      "[]()/*." 
                      #\\)
=> "\\*\\*\\*\\(Example String\\)\\*\\*\\*"

;; also can escape a singe char...
 (escape-special-chars "richardt@@vzavenue.net"
                      #\@@
                      #\@@)
=> "richardt@@@@vzavenue.net"
@end lisp CR279):;8< 	       hÈ   ·  - . , 3 #  	P#   #  4 544	
554$  "  5$  "  !44	
5$  "  5$   C 45$  4$  "  5"  6¯      g  str
	 È g  width	 È g  chr		 È g  rchr		 È g  len		0 È g  lpad		G È g  rpad	  È  g  filenamef  texinfo/string-utils.scm
 
	* ®		0 ®		3 ¯		6 ¯		= ¯	+	@ ¯	!	C ¯		G ¯		G ®		J ²		R ±		^ ²		b ²		i ´		l ´	 	s ´	/	v ´	%	y ´	 	 ±	  ´	  ®	  µ	  µ	   ·	) § ·	/ © ·	) ­ ·	% ® ·	> ¶ ±	 À ·	> Æ ·	Q È ·	 "	 È	  g  nameg  center-stringg  documentationf Returns a copy of @var{str} centered in a field of @var{width}
characters.  Any needed padding is done by character @var{chr}, which
defaults to @samp{#\space}.  If @var{rchr} is provided, then the
padding to the right will use it instead.  See the examples below.
left and @var{rchr} on the right.  The default @var{width} is 80.  The
default @var{chr} and @var{rchr} is @samp{#\space}.  The string is
never truncated.
@lisp
 (center-string "Richard Todd" 24)
=> "      Richard Todd      "

 (center-string " Richard Todd " 24 #\=)
=> "===== Richard Todd ====="

 (center-string " Richard Todd " 24 #\< #\>)
=> "<<<<< Richard Todd >>>>>"
@end lisp CR279:  hX   H  - . , 3 #  	P#   4 544
55$   C 6       @      g  str
		Q g  width		Q g  chr			Q g  len		'	Q g  pad		;	Q  g  filenamef  texinfo/string-utils.scm
 ¹
	! ¿		' ¿		* À		- À		4 À	 	7 À		; À		; ¿		B Á		F Á		Q Ã	 		Q	  g  nameg  left-justify-stringg  documentationf @code{left-justify-string str [width chr]}.  
Returns a copy of @var{str} padded with @var{chr} such that it is left
justified in a field of @var{width} characters.  The default
@var{width} is 80.  Unlike @samp{string-pad} from srfi-13, the string
is never truncated. CR279: hX   F  - . , 3 #  	P#   4 544
55$   C 6       >      g  str
		Q g  width		Q g  chr			Q g  len		'	Q g  pad		;	Q  g  filenamef  texinfo/string-utils.scm
 Å
	! Ê		' Ê		* Ë		- Ë		4 Ë	 	7 Ë		; Ë		; Ê		B Ì		F Ì		Q Î	 		Q	  g  nameg  right-justify-stringg  documentationf Returns a copy of @var{str} padded with @var{chr} such that it is
right justified in a field of @var{width} characters.  The default
@var{width} is 80.  The default @var{chr} is @samp{#\space}.  Unlike
@samp{string-pad} from srfi-13, the string is never truncated. CR))      h@   ¶   ]"  	
N NC4 M5$  4ML 5$  MNMLC"ÿÿÎ"ÿÿÊ    ®       g  c
		<  g  filenamef  texinfo/string-utils.scm
 ä			 í		 î		 å		 å		 å		 æ		) å		, é	&	. é		3 ê	 		<   C<   hX   Û  - . , 3 #   #  4 5$  A"   
HHO  6Ó      g  str
		X g  chr		X g  num			X g  prev-chr		5	L g  match-count		5	L g  repeat-locator		L	X  g  filenamef  texinfo/string-utils.scm
 Ð	  â		, â		5 â	
	L à		V ó	*	X ó	 		X	  g  nameg  collapse-repeated-charsg  documentationf  Returns a copy of @var{str} with all repeated instances of 
@var{chr} collapsed down to at most @var{num} instances.
The default value for @var{chr} is @samp{#\space}, and 
the default value for @var{num} is 1.

@lisp
 (collapse-repeated-chars "H  e  l  l  o")
=> "H e l l o"
 (collapse-repeated-chars "H--e--l--l--o" #\-)
=> "H-e-l-l-o"
 (collapse-repeated-chars "H-e--l---l----o" #\- 2)
=> "H-e--l--l--o"
@end lisp CR=>,3?    hp     ])45"  T4 5$  ;4 5$  4 5"ÿÿÂ4 566
"ÿÿ¢         g  str
		m g  non-wschars			m g  ans			c g  index			c g  next-non-ws			c g  next-ws		-	]  g  filenamef  texinfo/string-utils.scm
 ÷
	 ø			 ø		 ù		 û		 û		" ü		# þ		- þ	
	5 ÿ		6		C		M		P		[		]		c	
	c ù		d ù		m ù	 		m  g  nameg  split-by-single-words C@R2A h@   T  ]	4 5$  $.4 5$  .4 	5CCC       L      g  str
		9 g  len			9  g  filenamef  texinfo/string-utils.scm

	
			
										#			 			$		'		0	(	2		3		4		 		9  g  nameg  end-of-sentence?g  documentationf  6Return #t when STR likely denotes the end of sentence. CBRQ<)   h   e   ] 
6      ]       g  c
		
  g  filenamef  texinfo/string-utils.scm
E	$	
E	0 		
   CR?2:BS3@    hx    ]1 H 4J  5K L$  4J L5K "   L$  4J 5K "   4J 5K " (  
$  	"  6L45
$  45"  45$  Q"  45"  ,L$  "4	5$  4
5"  "ÿÿÏ"  "ÿÿÇ"ÿÿ_
$  L
"ÿÿAL$  044
5545L
"ÿÿ45L
"ÿþí4J 5L 
"ÿþÖ            g  str
	r g  ans	H[ g  words		H[ g  line		H[ g  count		H[ g  length-left	 [ g  	next-word	 [  g  filenamef  texinfo/string-utils.scm
C		E		E		H		I		$I		.L		/M		7M		<P		DP		HS		NW		TY		XY		]Z		eY	
	h_	 	o^		s`		w`		xa		}a	,	a	 b	 ^	
 e	 e	 c	 h	 ¢m	 ®i	 µj	 ¿i	 Àl	 Æl	, Êl	 Ýn	 ég	 íq	 ñc	 ös	s	c	|	|	.|	!|	"~	-	.~	=|	>	I	L	[	[S	\S	]T	rS	 =	r   C      hp   |  -  /     0   3  #  	P #  #  	#  #  #  #   O C      t      g  
line-width
		j g  expand-tabs?		j g  	tab-width			j g  collapse-whitespace?			j g  subsequent-indent			j g  initial-indent			j g  break-long-words?			j  g  filenamef  texinfo/string-utils.scm

	A	/	K	, 		j

g  
line-widthS
g  expand-tabs?Sg  	tab-widthS	g  collapse-whitespace?S	g  subsequent-indentS	g  initial-indentS	g  break-long-words?S	   g  nameg  make-text-wrapperg  documentationf Returns a procedure that will split a string into lines according to the
given parameters.

@table @code
@item #:line-width
This is the target length used when deciding where to wrap lines.
Default is 80.

@item #:expand-tabs?
Boolean describing whether tabs in the input should be expanded. Default
is #t.

@item #:tab-width
If tabs are expanded, this will be the number of spaces to which they
expand. Default is 8.

@item #:collapse-whitespace?
Boolean describing whether the whitespace inside the existing text
should be removed or not.  Default is #t.

If text is already well-formatted, and is just being wrapped to fit in a
different width, then set this to @samp{#f}. This way, many common text
conventions (such as two spaces between sentences) can be preserved if
in the original text. If the input text spacing cannot be trusted, then
leave this setting at the default, and all repeated whitespace will be
collapsed down to a single space.

@item #:initial-indent
Defines a string that will be put in front of the first line of wrapped
text. Default is the empty string, ``''.

@item #:subsequent-indent
Defines a string that will be put in front of all lines of wrapped
text, except the first one.  Default is the empty string, ``''.

@item #:break-long-words?
If a single word is too big to fit on a line, this setting tells the
wrapper what to do.  Defaults to #t, which will break up long words.
When set to #f, the line will be allowed, even though it is longer
than the defined @code{#:line-width}.
@end table

The return value is a procedure of one argument, the input string, which
returns a list of strings, where each element of the list is one line. C R    h   ô  - 1 3 4? 6    ì      g  str
			 g  kwargs			  g  filenamef  texinfo/string-utils.scm

	
			 			
  g  nameg  string->wrapped-linesg  documentationf 8@code{string->wrapped-lines str keywds ...}. Wraps the text given in
string @var{str} according to the parameters provided in @var{keywds},
or the default setting if they are not given. Returns a list of strings
representing the formatted lines. Valid keyword arguments are discussed
in @code{make-text-wrapper}. C"RT"UV     h    Ä  - 1 3 4 ?6      ¼      g  str
			 g  kwargs			  g  filenamef  texinfo/string-utils.scm

								 			
  g  nameg  fill-stringg  documentationf  Wraps the text given in string @var{str} according to the parameters
provided in @var{kwargs}, or the default setting if they are not
given.  Returns a single string with the wrapped text.  Valid keyword
arguments are discussed in @code{make-text-wrapper}. C!RC       ½       g  m
		,  g  filenamef  texinfo/string-utils.scm		
	(
Ð	s
Ü	}
 
; ¹
ñ Å
$F Ð&g ÷
(
6O
8o
:o
 	:q
   C6 