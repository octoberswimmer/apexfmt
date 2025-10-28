// Code generated from ./ApexParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ApexParser
import "github.com/antlr4-go/antlr/v4"

// ApexParserListener is a complete listener for a parse tree produced by ApexParser.
type ApexParserListener interface {
	antlr.ParseTreeListener

	// EnterTriggerUnit is called when entering the triggerUnit production.
	EnterTriggerUnit(c *TriggerUnitContext)

	// EnterTriggerCase is called when entering the triggerCase production.
	EnterTriggerCase(c *TriggerCaseContext)

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterEnumDeclaration is called when entering the enumDeclaration production.
	EnterEnumDeclaration(c *EnumDeclarationContext)

	// EnterEnumConstants is called when entering the enumConstants production.
	EnterEnumConstants(c *EnumConstantsContext)

	// EnterInterfaceDeclaration is called when entering the interfaceDeclaration production.
	EnterInterfaceDeclaration(c *InterfaceDeclarationContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterInterfaceBody is called when entering the interfaceBody production.
	EnterInterfaceBody(c *InterfaceBodyContext)

	// EnterClassBodyDeclaration is called when entering the classBodyDeclaration production.
	EnterClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterMemberDeclaration is called when entering the memberDeclaration production.
	EnterMemberDeclaration(c *MemberDeclarationContext)

	// EnterMethodDeclaration is called when entering the methodDeclaration production.
	EnterMethodDeclaration(c *MethodDeclarationContext)

	// EnterConstructorDeclaration is called when entering the constructorDeclaration production.
	EnterConstructorDeclaration(c *ConstructorDeclarationContext)

	// EnterFieldDeclaration is called when entering the fieldDeclaration production.
	EnterFieldDeclaration(c *FieldDeclarationContext)

	// EnterPropertyDeclaration is called when entering the propertyDeclaration production.
	EnterPropertyDeclaration(c *PropertyDeclarationContext)

	// EnterInterfaceMethodDeclaration is called when entering the interfaceMethodDeclaration production.
	EnterInterfaceMethodDeclaration(c *InterfaceMethodDeclarationContext)

	// EnterVariableDeclarators is called when entering the variableDeclarators production.
	EnterVariableDeclarators(c *VariableDeclaratorsContext)

	// EnterVariableDeclarator is called when entering the variableDeclarator production.
	EnterVariableDeclarator(c *VariableDeclaratorContext)

	// EnterArrayInitializer is called when entering the arrayInitializer production.
	EnterArrayInitializer(c *ArrayInitializerContext)

	// EnterTypeRef is called when entering the typeRef production.
	EnterTypeRef(c *TypeRefContext)

	// EnterArraySubscripts is called when entering the arraySubscripts production.
	EnterArraySubscripts(c *ArraySubscriptsContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterTypeArguments is called when entering the typeArguments production.
	EnterTypeArguments(c *TypeArgumentsContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterElementValuePairs is called when entering the elementValuePairs production.
	EnterElementValuePairs(c *ElementValuePairsContext)

	// EnterDelimitedElementValuePair is called when entering the delimitedElementValuePair production.
	EnterDelimitedElementValuePair(c *DelimitedElementValuePairContext)

	// EnterElementValuePair is called when entering the elementValuePair production.
	EnterElementValuePair(c *ElementValuePairContext)

	// EnterElementValue is called when entering the elementValue production.
	EnterElementValue(c *ElementValueContext)

	// EnterElementValueArrayInitializer is called when entering the elementValueArrayInitializer production.
	EnterElementValueArrayInitializer(c *ElementValueArrayInitializerContext)

	// EnterTrailingComma is called when entering the trailingComma production.
	EnterTrailingComma(c *TrailingCommaContext)

	// EnterTriggerBlock is called when entering the triggerBlock production.
	EnterTriggerBlock(c *TriggerBlockContext)

	// EnterTriggerStatement is called when entering the triggerStatement production.
	EnterTriggerStatement(c *TriggerStatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterLocalVariableDeclarationStatement is called when entering the localVariableDeclarationStatement production.
	EnterLocalVariableDeclarationStatement(c *LocalVariableDeclarationStatementContext)

	// EnterLocalVariableDeclaration is called when entering the localVariableDeclaration production.
	EnterLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBlockMemberDeclaration is called when entering the blockMemberDeclaration production.
	EnterBlockMemberDeclaration(c *BlockMemberDeclarationContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterWhenControl is called when entering the whenControl production.
	EnterWhenControl(c *WhenControlContext)

	// EnterWhenValue is called when entering the whenValue production.
	EnterWhenValue(c *WhenValueContext)

	// EnterWhenLiteral is called when entering the whenLiteral production.
	EnterWhenLiteral(c *WhenLiteralContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterDoWhileStatement is called when entering the doWhileStatement production.
	EnterDoWhileStatement(c *DoWhileStatementContext)

	// EnterTryStatement is called when entering the tryStatement production.
	EnterTryStatement(c *TryStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterThrowStatement is called when entering the throwStatement production.
	EnterThrowStatement(c *ThrowStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterInsertStatement is called when entering the insertStatement production.
	EnterInsertStatement(c *InsertStatementContext)

	// EnterUpdateStatement is called when entering the updateStatement production.
	EnterUpdateStatement(c *UpdateStatementContext)

	// EnterDeleteStatement is called when entering the deleteStatement production.
	EnterDeleteStatement(c *DeleteStatementContext)

	// EnterUndeleteStatement is called when entering the undeleteStatement production.
	EnterUndeleteStatement(c *UndeleteStatementContext)

	// EnterUpsertStatement is called when entering the upsertStatement production.
	EnterUpsertStatement(c *UpsertStatementContext)

	// EnterMergeStatement is called when entering the mergeStatement production.
	EnterMergeStatement(c *MergeStatementContext)

	// EnterRunAsStatement is called when entering the runAsStatement production.
	EnterRunAsStatement(c *RunAsStatementContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterPropertyBlock is called when entering the propertyBlock production.
	EnterPropertyBlock(c *PropertyBlockContext)

	// EnterGetter is called when entering the getter production.
	EnterGetter(c *GetterContext)

	// EnterSetter is called when entering the setter production.
	EnterSetter(c *SetterContext)

	// EnterCatchClause is called when entering the catchClause production.
	EnterCatchClause(c *CatchClauseContext)

	// EnterFinallyBlock is called when entering the finallyBlock production.
	EnterFinallyBlock(c *FinallyBlockContext)

	// EnterForControl is called when entering the forControl production.
	EnterForControl(c *ForControlContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterEnhancedForControl is called when entering the enhancedForControl production.
	EnterEnhancedForControl(c *EnhancedForControlContext)

	// EnterForUpdate is called when entering the forUpdate production.
	EnterForUpdate(c *ForUpdateContext)

	// EnterParExpression is called when entering the parExpression production.
	EnterParExpression(c *ParExpressionContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterArth1Expression is called when entering the arth1Expression production.
	EnterArth1Expression(c *Arth1ExpressionContext)

	// EnterCoalExpression is called when entering the coalExpression production.
	EnterCoalExpression(c *CoalExpressionContext)

	// EnterDotExpression is called when entering the dotExpression production.
	EnterDotExpression(c *DotExpressionContext)

	// EnterBitOrExpression is called when entering the bitOrExpression production.
	EnterBitOrExpression(c *BitOrExpressionContext)

	// EnterArrayExpression is called when entering the arrayExpression production.
	EnterArrayExpression(c *ArrayExpressionContext)

	// EnterAssignExpression is called when entering the assignExpression production.
	EnterAssignExpression(c *AssignExpressionContext)

	// EnterMethodCallExpression is called when entering the methodCallExpression production.
	EnterMethodCallExpression(c *MethodCallExpressionContext)

	// EnterBitNotExpression is called when entering the bitNotExpression production.
	EnterBitNotExpression(c *BitNotExpressionContext)

	// EnterNewInstanceExpression is called when entering the newInstanceExpression production.
	EnterNewInstanceExpression(c *NewInstanceExpressionContext)

	// EnterArth2Expression is called when entering the arth2Expression production.
	EnterArth2Expression(c *Arth2ExpressionContext)

	// EnterLogAndExpression is called when entering the logAndExpression production.
	EnterLogAndExpression(c *LogAndExpressionContext)

	// EnterCastExpression is called when entering the castExpression production.
	EnterCastExpression(c *CastExpressionContext)

	// EnterBitAndExpression is called when entering the bitAndExpression production.
	EnterBitAndExpression(c *BitAndExpressionContext)

	// EnterCmpExpression is called when entering the cmpExpression production.
	EnterCmpExpression(c *CmpExpressionContext)

	// EnterBitExpression is called when entering the bitExpression production.
	EnterBitExpression(c *BitExpressionContext)

	// EnterLogOrExpression is called when entering the logOrExpression production.
	EnterLogOrExpression(c *LogOrExpressionContext)

	// EnterCondExpression is called when entering the condExpression production.
	EnterCondExpression(c *CondExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterPostOpExpression is called when entering the postOpExpression production.
	EnterPostOpExpression(c *PostOpExpressionContext)

	// EnterNegExpression is called when entering the negExpression production.
	EnterNegExpression(c *NegExpressionContext)

	// EnterPreOpExpression is called when entering the preOpExpression production.
	EnterPreOpExpression(c *PreOpExpressionContext)

	// EnterSubExpression is called when entering the subExpression production.
	EnterSubExpression(c *SubExpressionContext)

	// EnterInstanceOfExpression is called when entering the instanceOfExpression production.
	EnterInstanceOfExpression(c *InstanceOfExpressionContext)

	// EnterThisPrimary is called when entering the thisPrimary production.
	EnterThisPrimary(c *ThisPrimaryContext)

	// EnterSuperPrimary is called when entering the superPrimary production.
	EnterSuperPrimary(c *SuperPrimaryContext)

	// EnterLiteralPrimary is called when entering the literalPrimary production.
	EnterLiteralPrimary(c *LiteralPrimaryContext)

	// EnterTypeRefPrimary is called when entering the typeRefPrimary production.
	EnterTypeRefPrimary(c *TypeRefPrimaryContext)

	// EnterIdPrimary is called when entering the idPrimary production.
	EnterIdPrimary(c *IdPrimaryContext)

	// EnterSoqlPrimary is called when entering the soqlPrimary production.
	EnterSoqlPrimary(c *SoqlPrimaryContext)

	// EnterSoslPrimary is called when entering the soslPrimary production.
	EnterSoslPrimary(c *SoslPrimaryContext)

	// EnterMethodCall is called when entering the methodCall production.
	EnterMethodCall(c *MethodCallContext)

	// EnterDotMethodCall is called when entering the dotMethodCall production.
	EnterDotMethodCall(c *DotMethodCallContext)

	// EnterCreator is called when entering the creator production.
	EnterCreator(c *CreatorContext)

	// EnterCreatedName is called when entering the createdName production.
	EnterCreatedName(c *CreatedNameContext)

	// EnterIdCreatedNamePair is called when entering the idCreatedNamePair production.
	EnterIdCreatedNamePair(c *IdCreatedNamePairContext)

	// EnterNoRest is called when entering the noRest production.
	EnterNoRest(c *NoRestContext)

	// EnterClassCreatorRest is called when entering the classCreatorRest production.
	EnterClassCreatorRest(c *ClassCreatorRestContext)

	// EnterArrayCreatorRest is called when entering the arrayCreatorRest production.
	EnterArrayCreatorRest(c *ArrayCreatorRestContext)

	// EnterMapCreatorRest is called when entering the mapCreatorRest production.
	EnterMapCreatorRest(c *MapCreatorRestContext)

	// EnterMapCreatorRestPair is called when entering the mapCreatorRestPair production.
	EnterMapCreatorRestPair(c *MapCreatorRestPairContext)

	// EnterSetCreatorRest is called when entering the setCreatorRest production.
	EnterSetCreatorRest(c *SetCreatorRestContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterSoqlLiteral is called when entering the soqlLiteral production.
	EnterSoqlLiteral(c *SoqlLiteralContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterSubQuery is called when entering the subQuery production.
	EnterSubQuery(c *SubQueryContext)

	// EnterSelectList is called when entering the selectList production.
	EnterSelectList(c *SelectListContext)

	// EnterSelectEntry is called when entering the selectEntry production.
	EnterSelectEntry(c *SelectEntryContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterFromNameList is called when entering the fromNameList production.
	EnterFromNameList(c *FromNameListContext)

	// EnterFieldNameAlias is called when entering the fieldNameAlias production.
	EnterFieldNameAlias(c *FieldNameAliasContext)

	// EnterSubFieldList is called when entering the subFieldList production.
	EnterSubFieldList(c *SubFieldListContext)

	// EnterSubFieldEntry is called when entering the subFieldEntry production.
	EnterSubFieldEntry(c *SubFieldEntryContext)

	// EnterSoqlFieldsParameter is called when entering the soqlFieldsParameter production.
	EnterSoqlFieldsParameter(c *SoqlFieldsParameterContext)

	// EnterSoqlFunction is called when entering the soqlFunction production.
	EnterSoqlFunction(c *SoqlFunctionContext)

	// EnterDateFieldName is called when entering the dateFieldName production.
	EnterDateFieldName(c *DateFieldNameContext)

	// EnterLocationValue is called when entering the locationValue production.
	EnterLocationValue(c *LocationValueContext)

	// EnterCoordinateValue is called when entering the coordinateValue production.
	EnterCoordinateValue(c *CoordinateValueContext)

	// EnterTypeOf is called when entering the typeOf production.
	EnterTypeOf(c *TypeOfContext)

	// EnterWhenClause is called when entering the whenClause production.
	EnterWhenClause(c *WhenClauseContext)

	// EnterElseClause is called when entering the elseClause production.
	EnterElseClause(c *ElseClauseContext)

	// EnterFieldNameList is called when entering the fieldNameList production.
	EnterFieldNameList(c *FieldNameListContext)

	// EnterUsingScope is called when entering the usingScope production.
	EnterUsingScope(c *UsingScopeContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterLogicalExpression is called when entering the logicalExpression production.
	EnterLogicalExpression(c *LogicalExpressionContext)

	// EnterConditionalExpression is called when entering the conditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterFieldExpression is called when entering the fieldExpression production.
	EnterFieldExpression(c *FieldExpressionContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterNullValue is called when entering the nullValue production.
	EnterNullValue(c *NullValueContext)

	// EnterBooleanLiteralValue is called when entering the booleanLiteralValue production.
	EnterBooleanLiteralValue(c *BooleanLiteralValueContext)

	// EnterSignedNumberValue is called when entering the signedNumberValue production.
	EnterSignedNumberValue(c *SignedNumberValueContext)

	// EnterStringLiteralValue is called when entering the stringLiteralValue production.
	EnterStringLiteralValue(c *StringLiteralValueContext)

	// EnterDateLiteralValue is called when entering the dateLiteralValue production.
	EnterDateLiteralValue(c *DateLiteralValueContext)

	// EnterTimeLiteralValue is called when entering the timeLiteralValue production.
	EnterTimeLiteralValue(c *TimeLiteralValueContext)

	// EnterDateTimeLiteralValue is called when entering the dateTimeLiteralValue production.
	EnterDateTimeLiteralValue(c *DateTimeLiteralValueContext)

	// EnterDateFormulaValue is called when entering the dateFormulaValue production.
	EnterDateFormulaValue(c *DateFormulaValueContext)

	// EnterCurrencyValueValue is called when entering the currencyValueValue production.
	EnterCurrencyValueValue(c *CurrencyValueValueContext)

	// EnterSubQueryValue is called when entering the subQueryValue production.
	EnterSubQueryValue(c *SubQueryValueContext)

	// EnterValueListValue is called when entering the valueListValue production.
	EnterValueListValue(c *ValueListValueContext)

	// EnterBoundExpressionValue is called when entering the boundExpressionValue production.
	EnterBoundExpressionValue(c *BoundExpressionValueContext)

	// EnterValueList is called when entering the valueList production.
	EnterValueList(c *ValueListContext)

	// EnterCurrencyValue is called when entering the currencyValue production.
	EnterCurrencyValue(c *CurrencyValueContext)

	// EnterSignedNumber is called when entering the signedNumber production.
	EnterSignedNumber(c *SignedNumberContext)

	// EnterWithClause is called when entering the withClause production.
	EnterWithClause(c *WithClauseContext)

	// EnterFilteringExpression is called when entering the filteringExpression production.
	EnterFilteringExpression(c *FilteringExpressionContext)

	// EnterDataCategorySelection is called when entering the dataCategorySelection production.
	EnterDataCategorySelection(c *DataCategorySelectionContext)

	// EnterDataCategoryName is called when entering the dataCategoryName production.
	EnterDataCategoryName(c *DataCategoryNameContext)

	// EnterFilteringSelector is called when entering the filteringSelector production.
	EnterFilteringSelector(c *FilteringSelectorContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterFieldOrderList is called when entering the fieldOrderList production.
	EnterFieldOrderList(c *FieldOrderListContext)

	// EnterFieldOrder is called when entering the fieldOrder production.
	EnterFieldOrder(c *FieldOrderContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterOffsetClause is called when entering the offsetClause production.
	EnterOffsetClause(c *OffsetClauseContext)

	// EnterAllRowsClause is called when entering the allRowsClause production.
	EnterAllRowsClause(c *AllRowsClauseContext)

	// EnterForClauses is called when entering the forClauses production.
	EnterForClauses(c *ForClausesContext)

	// EnterForClause is called when entering the forClause production.
	EnterForClause(c *ForClauseContext)

	// EnterBoundExpression is called when entering the boundExpression production.
	EnterBoundExpression(c *BoundExpressionContext)

	// EnterDateFormula is called when entering the dateFormula production.
	EnterDateFormula(c *DateFormulaContext)

	// EnterSignedInteger is called when entering the signedInteger production.
	EnterSignedInteger(c *SignedIntegerContext)

	// EnterSoqlId is called when entering the soqlId production.
	EnterSoqlId(c *SoqlIdContext)

	// EnterSoslLiteral is called when entering the soslLiteral production.
	EnterSoslLiteral(c *SoslLiteralContext)

	// EnterSoslLiteralAlt is called when entering the soslLiteralAlt production.
	EnterSoslLiteralAlt(c *SoslLiteralAltContext)

	// EnterSoslClauses is called when entering the soslClauses production.
	EnterSoslClauses(c *SoslClausesContext)

	// EnterInSearchGroup is called when entering the inSearchGroup production.
	EnterInSearchGroup(c *InSearchGroupContext)

	// EnterReturningFieldSpecList is called when entering the returningFieldSpecList production.
	EnterReturningFieldSpecList(c *ReturningFieldSpecListContext)

	// EnterWithDivisionAssign is called when entering the withDivisionAssign production.
	EnterWithDivisionAssign(c *WithDivisionAssignContext)

	// EnterWithDataCategory is called when entering the withDataCategory production.
	EnterWithDataCategory(c *WithDataCategoryContext)

	// EnterWithSnippet is called when entering the withSnippet production.
	EnterWithSnippet(c *WithSnippetContext)

	// EnterWithNetworkIn is called when entering the withNetworkIn production.
	EnterWithNetworkIn(c *WithNetworkInContext)

	// EnterWithNetworkAssign is called when entering the withNetworkAssign production.
	EnterWithNetworkAssign(c *WithNetworkAssignContext)

	// EnterWithPricebookIdAssign is called when entering the withPricebookIdAssign production.
	EnterWithPricebookIdAssign(c *WithPricebookIdAssignContext)

	// EnterWithMetadataAssign is called when entering the withMetadataAssign production.
	EnterWithMetadataAssign(c *WithMetadataAssignContext)

	// EnterWithModeClause is called when entering the withModeClause production.
	EnterWithModeClause(c *WithModeClauseContext)

	// EnterUpdateListClause is called when entering the updateListClause production.
	EnterUpdateListClause(c *UpdateListClauseContext)

	// EnterSearchGroup is called when entering the searchGroup production.
	EnterSearchGroup(c *SearchGroupContext)

	// EnterFieldSpecList is called when entering the fieldSpecList production.
	EnterFieldSpecList(c *FieldSpecListContext)

	// EnterFieldSpec is called when entering the fieldSpec production.
	EnterFieldSpec(c *FieldSpecContext)

	// EnterFieldSpecClauses is called when entering the fieldSpecClauses production.
	EnterFieldSpecClauses(c *FieldSpecClausesContext)

	// EnterFieldList is called when entering the fieldList production.
	EnterFieldList(c *FieldListContext)

	// EnterUpdateList is called when entering the updateList production.
	EnterUpdateList(c *UpdateListContext)

	// EnterUpdateType is called when entering the updateType production.
	EnterUpdateType(c *UpdateTypeContext)

	// EnterNetworkList is called when entering the networkList production.
	EnterNetworkList(c *NetworkListContext)

	// EnterSoslId is called when entering the soslId production.
	EnterSoslId(c *SoslIdContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterAnyId is called when entering the anyId production.
	EnterAnyId(c *AnyIdContext)

	// ExitTriggerUnit is called when exiting the triggerUnit production.
	ExitTriggerUnit(c *TriggerUnitContext)

	// ExitTriggerCase is called when exiting the triggerCase production.
	ExitTriggerCase(c *TriggerCaseContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitEnumDeclaration is called when exiting the enumDeclaration production.
	ExitEnumDeclaration(c *EnumDeclarationContext)

	// ExitEnumConstants is called when exiting the enumConstants production.
	ExitEnumConstants(c *EnumConstantsContext)

	// ExitInterfaceDeclaration is called when exiting the interfaceDeclaration production.
	ExitInterfaceDeclaration(c *InterfaceDeclarationContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitInterfaceBody is called when exiting the interfaceBody production.
	ExitInterfaceBody(c *InterfaceBodyContext)

	// ExitClassBodyDeclaration is called when exiting the classBodyDeclaration production.
	ExitClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitMemberDeclaration is called when exiting the memberDeclaration production.
	ExitMemberDeclaration(c *MemberDeclarationContext)

	// ExitMethodDeclaration is called when exiting the methodDeclaration production.
	ExitMethodDeclaration(c *MethodDeclarationContext)

	// ExitConstructorDeclaration is called when exiting the constructorDeclaration production.
	ExitConstructorDeclaration(c *ConstructorDeclarationContext)

	// ExitFieldDeclaration is called when exiting the fieldDeclaration production.
	ExitFieldDeclaration(c *FieldDeclarationContext)

	// ExitPropertyDeclaration is called when exiting the propertyDeclaration production.
	ExitPropertyDeclaration(c *PropertyDeclarationContext)

	// ExitInterfaceMethodDeclaration is called when exiting the interfaceMethodDeclaration production.
	ExitInterfaceMethodDeclaration(c *InterfaceMethodDeclarationContext)

	// ExitVariableDeclarators is called when exiting the variableDeclarators production.
	ExitVariableDeclarators(c *VariableDeclaratorsContext)

	// ExitVariableDeclarator is called when exiting the variableDeclarator production.
	ExitVariableDeclarator(c *VariableDeclaratorContext)

	// ExitArrayInitializer is called when exiting the arrayInitializer production.
	ExitArrayInitializer(c *ArrayInitializerContext)

	// ExitTypeRef is called when exiting the typeRef production.
	ExitTypeRef(c *TypeRefContext)

	// ExitArraySubscripts is called when exiting the arraySubscripts production.
	ExitArraySubscripts(c *ArraySubscriptsContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitTypeArguments is called when exiting the typeArguments production.
	ExitTypeArguments(c *TypeArgumentsContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitElementValuePairs is called when exiting the elementValuePairs production.
	ExitElementValuePairs(c *ElementValuePairsContext)

	// ExitDelimitedElementValuePair is called when exiting the delimitedElementValuePair production.
	ExitDelimitedElementValuePair(c *DelimitedElementValuePairContext)

	// ExitElementValuePair is called when exiting the elementValuePair production.
	ExitElementValuePair(c *ElementValuePairContext)

	// ExitElementValue is called when exiting the elementValue production.
	ExitElementValue(c *ElementValueContext)

	// ExitElementValueArrayInitializer is called when exiting the elementValueArrayInitializer production.
	ExitElementValueArrayInitializer(c *ElementValueArrayInitializerContext)

	// ExitTrailingComma is called when exiting the trailingComma production.
	ExitTrailingComma(c *TrailingCommaContext)

	// ExitTriggerBlock is called when exiting the triggerBlock production.
	ExitTriggerBlock(c *TriggerBlockContext)

	// ExitTriggerStatement is called when exiting the triggerStatement production.
	ExitTriggerStatement(c *TriggerStatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitLocalVariableDeclarationStatement is called when exiting the localVariableDeclarationStatement production.
	ExitLocalVariableDeclarationStatement(c *LocalVariableDeclarationStatementContext)

	// ExitLocalVariableDeclaration is called when exiting the localVariableDeclaration production.
	ExitLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBlockMemberDeclaration is called when exiting the blockMemberDeclaration production.
	ExitBlockMemberDeclaration(c *BlockMemberDeclarationContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitWhenControl is called when exiting the whenControl production.
	ExitWhenControl(c *WhenControlContext)

	// ExitWhenValue is called when exiting the whenValue production.
	ExitWhenValue(c *WhenValueContext)

	// ExitWhenLiteral is called when exiting the whenLiteral production.
	ExitWhenLiteral(c *WhenLiteralContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitDoWhileStatement is called when exiting the doWhileStatement production.
	ExitDoWhileStatement(c *DoWhileStatementContext)

	// ExitTryStatement is called when exiting the tryStatement production.
	ExitTryStatement(c *TryStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitThrowStatement is called when exiting the throwStatement production.
	ExitThrowStatement(c *ThrowStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitInsertStatement is called when exiting the insertStatement production.
	ExitInsertStatement(c *InsertStatementContext)

	// ExitUpdateStatement is called when exiting the updateStatement production.
	ExitUpdateStatement(c *UpdateStatementContext)

	// ExitDeleteStatement is called when exiting the deleteStatement production.
	ExitDeleteStatement(c *DeleteStatementContext)

	// ExitUndeleteStatement is called when exiting the undeleteStatement production.
	ExitUndeleteStatement(c *UndeleteStatementContext)

	// ExitUpsertStatement is called when exiting the upsertStatement production.
	ExitUpsertStatement(c *UpsertStatementContext)

	// ExitMergeStatement is called when exiting the mergeStatement production.
	ExitMergeStatement(c *MergeStatementContext)

	// ExitRunAsStatement is called when exiting the runAsStatement production.
	ExitRunAsStatement(c *RunAsStatementContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitPropertyBlock is called when exiting the propertyBlock production.
	ExitPropertyBlock(c *PropertyBlockContext)

	// ExitGetter is called when exiting the getter production.
	ExitGetter(c *GetterContext)

	// ExitSetter is called when exiting the setter production.
	ExitSetter(c *SetterContext)

	// ExitCatchClause is called when exiting the catchClause production.
	ExitCatchClause(c *CatchClauseContext)

	// ExitFinallyBlock is called when exiting the finallyBlock production.
	ExitFinallyBlock(c *FinallyBlockContext)

	// ExitForControl is called when exiting the forControl production.
	ExitForControl(c *ForControlContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitEnhancedForControl is called when exiting the enhancedForControl production.
	ExitEnhancedForControl(c *EnhancedForControlContext)

	// ExitForUpdate is called when exiting the forUpdate production.
	ExitForUpdate(c *ForUpdateContext)

	// ExitParExpression is called when exiting the parExpression production.
	ExitParExpression(c *ParExpressionContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitArth1Expression is called when exiting the arth1Expression production.
	ExitArth1Expression(c *Arth1ExpressionContext)

	// ExitCoalExpression is called when exiting the coalExpression production.
	ExitCoalExpression(c *CoalExpressionContext)

	// ExitDotExpression is called when exiting the dotExpression production.
	ExitDotExpression(c *DotExpressionContext)

	// ExitBitOrExpression is called when exiting the bitOrExpression production.
	ExitBitOrExpression(c *BitOrExpressionContext)

	// ExitArrayExpression is called when exiting the arrayExpression production.
	ExitArrayExpression(c *ArrayExpressionContext)

	// ExitAssignExpression is called when exiting the assignExpression production.
	ExitAssignExpression(c *AssignExpressionContext)

	// ExitMethodCallExpression is called when exiting the methodCallExpression production.
	ExitMethodCallExpression(c *MethodCallExpressionContext)

	// ExitBitNotExpression is called when exiting the bitNotExpression production.
	ExitBitNotExpression(c *BitNotExpressionContext)

	// ExitNewInstanceExpression is called when exiting the newInstanceExpression production.
	ExitNewInstanceExpression(c *NewInstanceExpressionContext)

	// ExitArth2Expression is called when exiting the arth2Expression production.
	ExitArth2Expression(c *Arth2ExpressionContext)

	// ExitLogAndExpression is called when exiting the logAndExpression production.
	ExitLogAndExpression(c *LogAndExpressionContext)

	// ExitCastExpression is called when exiting the castExpression production.
	ExitCastExpression(c *CastExpressionContext)

	// ExitBitAndExpression is called when exiting the bitAndExpression production.
	ExitBitAndExpression(c *BitAndExpressionContext)

	// ExitCmpExpression is called when exiting the cmpExpression production.
	ExitCmpExpression(c *CmpExpressionContext)

	// ExitBitExpression is called when exiting the bitExpression production.
	ExitBitExpression(c *BitExpressionContext)

	// ExitLogOrExpression is called when exiting the logOrExpression production.
	ExitLogOrExpression(c *LogOrExpressionContext)

	// ExitCondExpression is called when exiting the condExpression production.
	ExitCondExpression(c *CondExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitPostOpExpression is called when exiting the postOpExpression production.
	ExitPostOpExpression(c *PostOpExpressionContext)

	// ExitNegExpression is called when exiting the negExpression production.
	ExitNegExpression(c *NegExpressionContext)

	// ExitPreOpExpression is called when exiting the preOpExpression production.
	ExitPreOpExpression(c *PreOpExpressionContext)

	// ExitSubExpression is called when exiting the subExpression production.
	ExitSubExpression(c *SubExpressionContext)

	// ExitInstanceOfExpression is called when exiting the instanceOfExpression production.
	ExitInstanceOfExpression(c *InstanceOfExpressionContext)

	// ExitThisPrimary is called when exiting the thisPrimary production.
	ExitThisPrimary(c *ThisPrimaryContext)

	// ExitSuperPrimary is called when exiting the superPrimary production.
	ExitSuperPrimary(c *SuperPrimaryContext)

	// ExitLiteralPrimary is called when exiting the literalPrimary production.
	ExitLiteralPrimary(c *LiteralPrimaryContext)

	// ExitTypeRefPrimary is called when exiting the typeRefPrimary production.
	ExitTypeRefPrimary(c *TypeRefPrimaryContext)

	// ExitIdPrimary is called when exiting the idPrimary production.
	ExitIdPrimary(c *IdPrimaryContext)

	// ExitSoqlPrimary is called when exiting the soqlPrimary production.
	ExitSoqlPrimary(c *SoqlPrimaryContext)

	// ExitSoslPrimary is called when exiting the soslPrimary production.
	ExitSoslPrimary(c *SoslPrimaryContext)

	// ExitMethodCall is called when exiting the methodCall production.
	ExitMethodCall(c *MethodCallContext)

	// ExitDotMethodCall is called when exiting the dotMethodCall production.
	ExitDotMethodCall(c *DotMethodCallContext)

	// ExitCreator is called when exiting the creator production.
	ExitCreator(c *CreatorContext)

	// ExitCreatedName is called when exiting the createdName production.
	ExitCreatedName(c *CreatedNameContext)

	// ExitIdCreatedNamePair is called when exiting the idCreatedNamePair production.
	ExitIdCreatedNamePair(c *IdCreatedNamePairContext)

	// ExitNoRest is called when exiting the noRest production.
	ExitNoRest(c *NoRestContext)

	// ExitClassCreatorRest is called when exiting the classCreatorRest production.
	ExitClassCreatorRest(c *ClassCreatorRestContext)

	// ExitArrayCreatorRest is called when exiting the arrayCreatorRest production.
	ExitArrayCreatorRest(c *ArrayCreatorRestContext)

	// ExitMapCreatorRest is called when exiting the mapCreatorRest production.
	ExitMapCreatorRest(c *MapCreatorRestContext)

	// ExitMapCreatorRestPair is called when exiting the mapCreatorRestPair production.
	ExitMapCreatorRestPair(c *MapCreatorRestPairContext)

	// ExitSetCreatorRest is called when exiting the setCreatorRest production.
	ExitSetCreatorRest(c *SetCreatorRestContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitSoqlLiteral is called when exiting the soqlLiteral production.
	ExitSoqlLiteral(c *SoqlLiteralContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitSubQuery is called when exiting the subQuery production.
	ExitSubQuery(c *SubQueryContext)

	// ExitSelectList is called when exiting the selectList production.
	ExitSelectList(c *SelectListContext)

	// ExitSelectEntry is called when exiting the selectEntry production.
	ExitSelectEntry(c *SelectEntryContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitFromNameList is called when exiting the fromNameList production.
	ExitFromNameList(c *FromNameListContext)

	// ExitFieldNameAlias is called when exiting the fieldNameAlias production.
	ExitFieldNameAlias(c *FieldNameAliasContext)

	// ExitSubFieldList is called when exiting the subFieldList production.
	ExitSubFieldList(c *SubFieldListContext)

	// ExitSubFieldEntry is called when exiting the subFieldEntry production.
	ExitSubFieldEntry(c *SubFieldEntryContext)

	// ExitSoqlFieldsParameter is called when exiting the soqlFieldsParameter production.
	ExitSoqlFieldsParameter(c *SoqlFieldsParameterContext)

	// ExitSoqlFunction is called when exiting the soqlFunction production.
	ExitSoqlFunction(c *SoqlFunctionContext)

	// ExitDateFieldName is called when exiting the dateFieldName production.
	ExitDateFieldName(c *DateFieldNameContext)

	// ExitLocationValue is called when exiting the locationValue production.
	ExitLocationValue(c *LocationValueContext)

	// ExitCoordinateValue is called when exiting the coordinateValue production.
	ExitCoordinateValue(c *CoordinateValueContext)

	// ExitTypeOf is called when exiting the typeOf production.
	ExitTypeOf(c *TypeOfContext)

	// ExitWhenClause is called when exiting the whenClause production.
	ExitWhenClause(c *WhenClauseContext)

	// ExitElseClause is called when exiting the elseClause production.
	ExitElseClause(c *ElseClauseContext)

	// ExitFieldNameList is called when exiting the fieldNameList production.
	ExitFieldNameList(c *FieldNameListContext)

	// ExitUsingScope is called when exiting the usingScope production.
	ExitUsingScope(c *UsingScopeContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitLogicalExpression is called when exiting the logicalExpression production.
	ExitLogicalExpression(c *LogicalExpressionContext)

	// ExitConditionalExpression is called when exiting the conditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitFieldExpression is called when exiting the fieldExpression production.
	ExitFieldExpression(c *FieldExpressionContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitNullValue is called when exiting the nullValue production.
	ExitNullValue(c *NullValueContext)

	// ExitBooleanLiteralValue is called when exiting the booleanLiteralValue production.
	ExitBooleanLiteralValue(c *BooleanLiteralValueContext)

	// ExitSignedNumberValue is called when exiting the signedNumberValue production.
	ExitSignedNumberValue(c *SignedNumberValueContext)

	// ExitStringLiteralValue is called when exiting the stringLiteralValue production.
	ExitStringLiteralValue(c *StringLiteralValueContext)

	// ExitDateLiteralValue is called when exiting the dateLiteralValue production.
	ExitDateLiteralValue(c *DateLiteralValueContext)

	// ExitTimeLiteralValue is called when exiting the timeLiteralValue production.
	ExitTimeLiteralValue(c *TimeLiteralValueContext)

	// ExitDateTimeLiteralValue is called when exiting the dateTimeLiteralValue production.
	ExitDateTimeLiteralValue(c *DateTimeLiteralValueContext)

	// ExitDateFormulaValue is called when exiting the dateFormulaValue production.
	ExitDateFormulaValue(c *DateFormulaValueContext)

	// ExitCurrencyValueValue is called when exiting the currencyValueValue production.
	ExitCurrencyValueValue(c *CurrencyValueValueContext)

	// ExitSubQueryValue is called when exiting the subQueryValue production.
	ExitSubQueryValue(c *SubQueryValueContext)

	// ExitValueListValue is called when exiting the valueListValue production.
	ExitValueListValue(c *ValueListValueContext)

	// ExitBoundExpressionValue is called when exiting the boundExpressionValue production.
	ExitBoundExpressionValue(c *BoundExpressionValueContext)

	// ExitValueList is called when exiting the valueList production.
	ExitValueList(c *ValueListContext)

	// ExitCurrencyValue is called when exiting the currencyValue production.
	ExitCurrencyValue(c *CurrencyValueContext)

	// ExitSignedNumber is called when exiting the signedNumber production.
	ExitSignedNumber(c *SignedNumberContext)

	// ExitWithClause is called when exiting the withClause production.
	ExitWithClause(c *WithClauseContext)

	// ExitFilteringExpression is called when exiting the filteringExpression production.
	ExitFilteringExpression(c *FilteringExpressionContext)

	// ExitDataCategorySelection is called when exiting the dataCategorySelection production.
	ExitDataCategorySelection(c *DataCategorySelectionContext)

	// ExitDataCategoryName is called when exiting the dataCategoryName production.
	ExitDataCategoryName(c *DataCategoryNameContext)

	// ExitFilteringSelector is called when exiting the filteringSelector production.
	ExitFilteringSelector(c *FilteringSelectorContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitFieldOrderList is called when exiting the fieldOrderList production.
	ExitFieldOrderList(c *FieldOrderListContext)

	// ExitFieldOrder is called when exiting the fieldOrder production.
	ExitFieldOrder(c *FieldOrderContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitOffsetClause is called when exiting the offsetClause production.
	ExitOffsetClause(c *OffsetClauseContext)

	// ExitAllRowsClause is called when exiting the allRowsClause production.
	ExitAllRowsClause(c *AllRowsClauseContext)

	// ExitForClauses is called when exiting the forClauses production.
	ExitForClauses(c *ForClausesContext)

	// ExitForClause is called when exiting the forClause production.
	ExitForClause(c *ForClauseContext)

	// ExitBoundExpression is called when exiting the boundExpression production.
	ExitBoundExpression(c *BoundExpressionContext)

	// ExitDateFormula is called when exiting the dateFormula production.
	ExitDateFormula(c *DateFormulaContext)

	// ExitSignedInteger is called when exiting the signedInteger production.
	ExitSignedInteger(c *SignedIntegerContext)

	// ExitSoqlId is called when exiting the soqlId production.
	ExitSoqlId(c *SoqlIdContext)

	// ExitSoslLiteral is called when exiting the soslLiteral production.
	ExitSoslLiteral(c *SoslLiteralContext)

	// ExitSoslLiteralAlt is called when exiting the soslLiteralAlt production.
	ExitSoslLiteralAlt(c *SoslLiteralAltContext)

	// ExitSoslClauses is called when exiting the soslClauses production.
	ExitSoslClauses(c *SoslClausesContext)

	// ExitInSearchGroup is called when exiting the inSearchGroup production.
	ExitInSearchGroup(c *InSearchGroupContext)

	// ExitReturningFieldSpecList is called when exiting the returningFieldSpecList production.
	ExitReturningFieldSpecList(c *ReturningFieldSpecListContext)

	// ExitWithDivisionAssign is called when exiting the withDivisionAssign production.
	ExitWithDivisionAssign(c *WithDivisionAssignContext)

	// ExitWithDataCategory is called when exiting the withDataCategory production.
	ExitWithDataCategory(c *WithDataCategoryContext)

	// ExitWithSnippet is called when exiting the withSnippet production.
	ExitWithSnippet(c *WithSnippetContext)

	// ExitWithNetworkIn is called when exiting the withNetworkIn production.
	ExitWithNetworkIn(c *WithNetworkInContext)

	// ExitWithNetworkAssign is called when exiting the withNetworkAssign production.
	ExitWithNetworkAssign(c *WithNetworkAssignContext)

	// ExitWithPricebookIdAssign is called when exiting the withPricebookIdAssign production.
	ExitWithPricebookIdAssign(c *WithPricebookIdAssignContext)

	// ExitWithMetadataAssign is called when exiting the withMetadataAssign production.
	ExitWithMetadataAssign(c *WithMetadataAssignContext)

	// ExitWithModeClause is called when exiting the withModeClause production.
	ExitWithModeClause(c *WithModeClauseContext)

	// ExitUpdateListClause is called when exiting the updateListClause production.
	ExitUpdateListClause(c *UpdateListClauseContext)

	// ExitSearchGroup is called when exiting the searchGroup production.
	ExitSearchGroup(c *SearchGroupContext)

	// ExitFieldSpecList is called when exiting the fieldSpecList production.
	ExitFieldSpecList(c *FieldSpecListContext)

	// ExitFieldSpec is called when exiting the fieldSpec production.
	ExitFieldSpec(c *FieldSpecContext)

	// ExitFieldSpecClauses is called when exiting the fieldSpecClauses production.
	ExitFieldSpecClauses(c *FieldSpecClausesContext)

	// ExitFieldList is called when exiting the fieldList production.
	ExitFieldList(c *FieldListContext)

	// ExitUpdateList is called when exiting the updateList production.
	ExitUpdateList(c *UpdateListContext)

	// ExitUpdateType is called when exiting the updateType production.
	ExitUpdateType(c *UpdateTypeContext)

	// ExitNetworkList is called when exiting the networkList production.
	ExitNetworkList(c *NetworkListContext)

	// ExitSoslId is called when exiting the soslId production.
	ExitSoslId(c *SoslIdContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitAnyId is called when exiting the anyId production.
	ExitAnyId(c *AnyIdContext)
}
