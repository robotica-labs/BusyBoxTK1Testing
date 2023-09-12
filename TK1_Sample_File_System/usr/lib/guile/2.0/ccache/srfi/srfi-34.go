GOOF----LE-4-2.0C      ] z 4    hC      ] g  guile�	 �	g  define-module*�	 �	 �	g  srfi�	g  srfi-34�	 �		g  filenameS�	
f  srfi/srfi-34.scm�	g  exportsS�	g  with-exception-handler�	g  guard�	 �	g  replacementsS�	g  raise�	 �	g  set-current-module�	 �	 �	g  cond-expand-provide�	g  current-module�	 �	g  	throw-key�	g  with-throw-handler�	g  throw�	g  make-syntax-transformer�	 �	 �	g  macro�	g  $sc-dispatch�	  �	! �	"g  any�	#g  each+�	$g  free-id�	%g  syntax-object�	&g  else�	'g  top�	(' �	)g  ribcage�	*) �	+g  x�	,+ �	-g  m-G5iZBnIeGOOhGY2SdwEyWq-1521�	.-' �	/. �	0f  l-G5iZBnIeGOOhGY2SdwEyWq-1523�	10 �	2),/1 �	3(*2 �	4g  hygiene�	54 �	6%&35 �	7$6 �	8g  each-any�	9"8��	:79��	;: �	<#"; �	="<��	>=9��	?">��	@g  catch�	Ag  dummy�	Bg  var�	Cg  clause�	Dg  e�	Eg  e*�	Fg  body�	Gg  body*�	HABCDEFG �	I.(((((( �	Jf  l-G5iZBnIeGOOhGY2SdwEyWq-1526�	Kf  l-G5iZBnIeGOOhGY2SdwEyWq-1527�	Lf  l-G5iZBnIeGOOhGY2SdwEyWq-1528�	Mf  l-G5iZBnIeGOOhGY2SdwEyWq-1529�	Nf  l-G5iZBnIeGOOhGY2SdwEyWq-1530�	Of  l-G5iZBnIeGOOhGY2SdwEyWq-1531�	Pf  l-G5iZBnIeGOOhGY2SdwEyWq-1532�	QJKLMNOP �	R)HIQ �	S(R*2 �	T%@S5 �	U%S5 �	Vg  lambda�	W%VS5 �	Xg  key�	Y%XS5 �	Zg  cond�	[%ZS5 �	\g  append�	]\ �	^\ �	_%&S5 �	`"9��	a`9��	b"a��	cg  clause*�	dABCcFG �	e.((((( �	ff  l-G5iZBnIeGOOhGY2SdwEyWq-1544�	gf  l-G5iZBnIeGOOhGY2SdwEyWq-1545�	hf  l-G5iZBnIeGOOhGY2SdwEyWq-1546�	if  l-G5iZBnIeGOOhGY2SdwEyWq-1547�	jf  l-G5iZBnIeGOOhGY2SdwEyWq-1548�	kf  l-G5iZBnIeGOOhGY2SdwEyWq-1549�	lfghijk �	m)del �	n(m*2 �	o%@n5 �	p%n5 �	q%Vn5 �	r%Xn5 �	s%Zn5 �	t%&n5 �	u%n5 �	vg  syntax-violation�	wv �	xv �	yf  -source expression failed to match any pattern�C 5h�
  �   ]4	
5 4 >  "  G   4i4i5 >  "  G  R     h   n   ]L 6f       g  key
		 g  obj		  g  filenamef  srfi/srfi-34.scm�
	,	��		-	�� 			   C       h   �  ] O 6       �      g  handler
		 g  thunk		  g  filenamef  srfi/srfi-34.scm�
	%
��		*	�� 			  g  nameg  with-exception-handler�g  documentationf  �Returns the result(s) of invoking THUNK. HANDLER must be a
procedure that accepts one argument.  It is installed as the current
exception handler for the dynamic extent (as determined by
dynamic-wind) of the invocation of THUNK.� CR     h   �  ] 6      �      g  obj
		
  g  filenamef  srfi/srfi-34.scm�
	/
��	
	5	�� 		
  g  nameg  raise�g  documentationf 6Invokes the current exception handler on OBJ.  The handler is
called in the dynamic environment of the call to raise, except that
the current exception handler is that in place for the call to
with-exception-handler that installed the handler being called.  The
handler's continuation is otherwise unspecified.� CR4!?TUWY[^_     h8   �   ]��� 4�� 5�  C   �       g  dummy
		5 g  var		5 g  clause			5 g  e			5 g  e*			5 g  body			5 g  body*			5  		5	   Cbopqrs^tu 	   h@   �   ]��� 4   5��  C  �       g  dummy
		> g  var		> g  clause			> g  clause*			> g  body			> g  body*			>  		>	   Cxy     h@     ]4 5$  @4 5$  @ 6             g  x
		9 g  tmp		9 g  tmp		"	9  g  filenamef  srfi/srfi-34.scm�
	8	�� 		9  g  documentationf �Syntax: (guard (<var> <clause1> <clause2> ...) <body>)
Each <clause> should have the same form as a `cond' clause.

Semantics: Evaluating a guard form evaluates <body> with an exception
handler that binds the raised object to <var> and within the scope of
that binding evaluates the clauses as if they were the clauses of a
cond expression.  That implicit cond expression is evaluated with the
continuation and dynamic environment of the guard expression.  If
every <clause>'s <test> evaluates to false and there is no else
clause, then raise is re-invoked on the raised object within the
dynamic environment of the original call to raise except that the
current exception handler is that of the guard expression.�g  
macro-typeg  syntax-rules�g  patternsg  varg  clauseg  ...g  elseg  eg  e*g  ...  g  bodyg  body*g  ... g  varg  clauseg  clause*g  ... g  bodyg  body*g  ...   C5RC    �       g  m
		,  g  filenamef  srfi/srfi-34.scm�		
��	-	!
��	2	!	��	:	!	&��	?	!
��	I	#	��	L	#
���	%
��u	/
�� 
	
�
   C6 