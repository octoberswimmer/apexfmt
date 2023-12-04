// Code generated from ./ApexParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ApexParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ApexParser.
type ApexParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ApexParser#triggerUnit.
	VisitTriggerUnit(ctx *TriggerUnitContext) interface{}

	// Visit a parse tree produced by ApexParser#triggerCase.
	VisitTriggerCase(ctx *TriggerCaseContext) interface{}

	// Visit a parse tree produced by ApexParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by ApexParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by ApexParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by ApexParser#enumDeclaration.
	VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{}

	// Visit a parse tree produced by ApexParser#enumConstants.
	VisitEnumConstants(ctx *EnumConstantsContext) interface{}

	// Visit a parse tree produced by ApexParser#interfaceDeclaration.
	VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by ApexParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by ApexParser#classBody.
	VisitClassBody(ctx *ClassBodyContext) interface{}

	// Visit a parse tree produced by ApexParser#interfaceBody.
	VisitInterfaceBody(ctx *InterfaceBodyContext) interface{}

	// Visit a parse tree produced by ApexParser#classBodyDeclaration.
	VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) interface{}

	// Visit a parse tree produced by ApexParser#modifier.
	VisitModifier(ctx *ModifierContext) interface{}

	// Visit a parse tree produced by ApexParser#memberDeclaration.
	VisitMemberDeclaration(ctx *MemberDeclarationContext) interface{}

	// Visit a parse tree produced by ApexParser#methodDeclaration.
	VisitMethodDeclaration(ctx *MethodDeclarationContext) interface{}

	// Visit a parse tree produced by ApexParser#constructorDeclaration.
	VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) interface{}

	// Visit a parse tree produced by ApexParser#fieldDeclaration.
	VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{}

	// Visit a parse tree produced by ApexParser#propertyDeclaration.
	VisitPropertyDeclaration(ctx *PropertyDeclarationContext) interface{}

	// Visit a parse tree produced by ApexParser#interfaceMethodDeclaration.
	VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) interface{}

	// Visit a parse tree produced by ApexParser#variableDeclarators.
	VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{}

	// Visit a parse tree produced by ApexParser#variableDeclarator.
	VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{}

	// Visit a parse tree produced by ApexParser#arrayInitializer.
	VisitArrayInitializer(ctx *ArrayInitializerContext) interface{}

	// Visit a parse tree produced by ApexParser#typeRef.
	VisitTypeRef(ctx *TypeRefContext) interface{}

	// Visit a parse tree produced by ApexParser#arraySubscripts.
	VisitArraySubscripts(ctx *ArraySubscriptsContext) interface{}

	// Visit a parse tree produced by ApexParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by ApexParser#typeArguments.
	VisitTypeArguments(ctx *TypeArgumentsContext) interface{}

	// Visit a parse tree produced by ApexParser#formalParameters.
	VisitFormalParameters(ctx *FormalParametersContext) interface{}

	// Visit a parse tree produced by ApexParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by ApexParser#formalParameter.
	VisitFormalParameter(ctx *FormalParameterContext) interface{}

	// Visit a parse tree produced by ApexParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}

	// Visit a parse tree produced by ApexParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by ApexParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by ApexParser#elementValuePairs.
	VisitElementValuePairs(ctx *ElementValuePairsContext) interface{}

	// Visit a parse tree produced by ApexParser#delimitedElementValuePair.
	VisitDelimitedElementValuePair(ctx *DelimitedElementValuePairContext) interface{}

	// Visit a parse tree produced by ApexParser#elementValuePair.
	VisitElementValuePair(ctx *ElementValuePairContext) interface{}

	// Visit a parse tree produced by ApexParser#elementValue.
	VisitElementValue(ctx *ElementValueContext) interface{}

	// Visit a parse tree produced by ApexParser#elementValueArrayInitializer.
	VisitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) interface{}

	// Visit a parse tree produced by ApexParser#trailingComma.
	VisitTrailingComma(ctx *TrailingCommaContext) interface{}

	// Visit a parse tree produced by ApexParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ApexParser#localVariableDeclarationStatement.
	VisitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#localVariableDeclaration.
	VisitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) interface{}

	// Visit a parse tree produced by ApexParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ApexParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#whenControl.
	VisitWhenControl(ctx *WhenControlContext) interface{}

	// Visit a parse tree produced by ApexParser#whenValue.
	VisitWhenValue(ctx *WhenValueContext) interface{}

	// Visit a parse tree produced by ApexParser#whenLiteral.
	VisitWhenLiteral(ctx *WhenLiteralContext) interface{}

	// Visit a parse tree produced by ApexParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#doWhileStatement.
	VisitDoWhileStatement(ctx *DoWhileStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#tryStatement.
	VisitTryStatement(ctx *TryStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#throwStatement.
	VisitThrowStatement(ctx *ThrowStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#insertStatement.
	VisitInsertStatement(ctx *InsertStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#updateStatement.
	VisitUpdateStatement(ctx *UpdateStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#undeleteStatement.
	VisitUndeleteStatement(ctx *UndeleteStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#upsertStatement.
	VisitUpsertStatement(ctx *UpsertStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#mergeStatement.
	VisitMergeStatement(ctx *MergeStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#runAsStatement.
	VisitRunAsStatement(ctx *RunAsStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by ApexParser#propertyBlock.
	VisitPropertyBlock(ctx *PropertyBlockContext) interface{}

	// Visit a parse tree produced by ApexParser#getter.
	VisitGetter(ctx *GetterContext) interface{}

	// Visit a parse tree produced by ApexParser#setter.
	VisitSetter(ctx *SetterContext) interface{}

	// Visit a parse tree produced by ApexParser#catchClause.
	VisitCatchClause(ctx *CatchClauseContext) interface{}

	// Visit a parse tree produced by ApexParser#finallyBlock.
	VisitFinallyBlock(ctx *FinallyBlockContext) interface{}

	// Visit a parse tree produced by ApexParser#forControl.
	VisitForControl(ctx *ForControlContext) interface{}

	// Visit a parse tree produced by ApexParser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by ApexParser#enhancedForControl.
	VisitEnhancedForControl(ctx *EnhancedForControlContext) interface{}

	// Visit a parse tree produced by ApexParser#forUpdate.
	VisitForUpdate(ctx *ForUpdateContext) interface{}

	// Visit a parse tree produced by ApexParser#parExpression.
	VisitParExpression(ctx *ParExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by ApexParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#arth1Expression.
	VisitArth1Expression(ctx *Arth1ExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#dotExpression.
	VisitDotExpression(ctx *DotExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#bitOrExpression.
	VisitBitOrExpression(ctx *BitOrExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#arrayExpression.
	VisitArrayExpression(ctx *ArrayExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#assignExpression.
	VisitAssignExpression(ctx *AssignExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#methodCallExpression.
	VisitMethodCallExpression(ctx *MethodCallExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#bitNotExpression.
	VisitBitNotExpression(ctx *BitNotExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#newInstanceExpression.
	VisitNewInstanceExpression(ctx *NewInstanceExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#arth2Expression.
	VisitArth2Expression(ctx *Arth2ExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#logAndExpression.
	VisitLogAndExpression(ctx *LogAndExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#castExpression.
	VisitCastExpression(ctx *CastExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#bitAndExpression.
	VisitBitAndExpression(ctx *BitAndExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#cmpExpression.
	VisitCmpExpression(ctx *CmpExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#bitExpression.
	VisitBitExpression(ctx *BitExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#logOrExpression.
	VisitLogOrExpression(ctx *LogOrExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#condExpression.
	VisitCondExpression(ctx *CondExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#postOpExpression.
	VisitPostOpExpression(ctx *PostOpExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#negExpression.
	VisitNegExpression(ctx *NegExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#preOpExpression.
	VisitPreOpExpression(ctx *PreOpExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#subExpression.
	VisitSubExpression(ctx *SubExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#instanceOfExpression.
	VisitInstanceOfExpression(ctx *InstanceOfExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#thisPrimary.
	VisitThisPrimary(ctx *ThisPrimaryContext) interface{}

	// Visit a parse tree produced by ApexParser#superPrimary.
	VisitSuperPrimary(ctx *SuperPrimaryContext) interface{}

	// Visit a parse tree produced by ApexParser#literalPrimary.
	VisitLiteralPrimary(ctx *LiteralPrimaryContext) interface{}

	// Visit a parse tree produced by ApexParser#typeRefPrimary.
	VisitTypeRefPrimary(ctx *TypeRefPrimaryContext) interface{}

	// Visit a parse tree produced by ApexParser#idPrimary.
	VisitIdPrimary(ctx *IdPrimaryContext) interface{}

	// Visit a parse tree produced by ApexParser#soqlPrimary.
	VisitSoqlPrimary(ctx *SoqlPrimaryContext) interface{}

	// Visit a parse tree produced by ApexParser#soslPrimary.
	VisitSoslPrimary(ctx *SoslPrimaryContext) interface{}

	// Visit a parse tree produced by ApexParser#methodCall.
	VisitMethodCall(ctx *MethodCallContext) interface{}

	// Visit a parse tree produced by ApexParser#dotMethodCall.
	VisitDotMethodCall(ctx *DotMethodCallContext) interface{}

	// Visit a parse tree produced by ApexParser#creator.
	VisitCreator(ctx *CreatorContext) interface{}

	// Visit a parse tree produced by ApexParser#createdName.
	VisitCreatedName(ctx *CreatedNameContext) interface{}

	// Visit a parse tree produced by ApexParser#idCreatedNamePair.
	VisitIdCreatedNamePair(ctx *IdCreatedNamePairContext) interface{}

	// Visit a parse tree produced by ApexParser#noRest.
	VisitNoRest(ctx *NoRestContext) interface{}

	// Visit a parse tree produced by ApexParser#classCreatorRest.
	VisitClassCreatorRest(ctx *ClassCreatorRestContext) interface{}

	// Visit a parse tree produced by ApexParser#arrayCreatorRest.
	VisitArrayCreatorRest(ctx *ArrayCreatorRestContext) interface{}

	// Visit a parse tree produced by ApexParser#mapCreatorRest.
	VisitMapCreatorRest(ctx *MapCreatorRestContext) interface{}

	// Visit a parse tree produced by ApexParser#mapCreatorRestPair.
	VisitMapCreatorRestPair(ctx *MapCreatorRestPairContext) interface{}

	// Visit a parse tree produced by ApexParser#setCreatorRest.
	VisitSetCreatorRest(ctx *SetCreatorRestContext) interface{}

	// Visit a parse tree produced by ApexParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by ApexParser#soqlLiteral.
	VisitSoqlLiteral(ctx *SoqlLiteralContext) interface{}

	// Visit a parse tree produced by ApexParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by ApexParser#subQuery.
	VisitSubQuery(ctx *SubQueryContext) interface{}

	// Visit a parse tree produced by ApexParser#selectList.
	VisitSelectList(ctx *SelectListContext) interface{}

	// Visit a parse tree produced by ApexParser#selectEntry.
	VisitSelectEntry(ctx *SelectEntryContext) interface{}

	// Visit a parse tree produced by ApexParser#fieldName.
	VisitFieldName(ctx *FieldNameContext) interface{}

	// Visit a parse tree produced by ApexParser#fromNameList.
	VisitFromNameList(ctx *FromNameListContext) interface{}

	// Visit a parse tree produced by ApexParser#fieldNameAlias.
	VisitFieldNameAlias(ctx *FieldNameAliasContext) interface{}

	// Visit a parse tree produced by ApexParser#subFieldList.
	VisitSubFieldList(ctx *SubFieldListContext) interface{}

	// Visit a parse tree produced by ApexParser#subFieldEntry.
	VisitSubFieldEntry(ctx *SubFieldEntryContext) interface{}

	// Visit a parse tree produced by ApexParser#soqlFieldsParameter.
	VisitSoqlFieldsParameter(ctx *SoqlFieldsParameterContext) interface{}

	// Visit a parse tree produced by ApexParser#soqlFunction.
	VisitSoqlFunction(ctx *SoqlFunctionContext) interface{}

	// Visit a parse tree produced by ApexParser#dateFieldName.
	VisitDateFieldName(ctx *DateFieldNameContext) interface{}

	// Visit a parse tree produced by ApexParser#typeOf.
	VisitTypeOf(ctx *TypeOfContext) interface{}

	// Visit a parse tree produced by ApexParser#whenClause.
	VisitWhenClause(ctx *WhenClauseContext) interface{}

	// Visit a parse tree produced by ApexParser#elseClause.
	VisitElseClause(ctx *ElseClauseContext) interface{}

	// Visit a parse tree produced by ApexParser#fieldNameList.
	VisitFieldNameList(ctx *FieldNameListContext) interface{}

	// Visit a parse tree produced by ApexParser#usingScope.
	VisitUsingScope(ctx *UsingScopeContext) interface{}

	// Visit a parse tree produced by ApexParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by ApexParser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#conditionalExpression.
	VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#fieldExpression.
	VisitFieldExpression(ctx *FieldExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by ApexParser#nullValue.
	VisitNullValue(ctx *NullValueContext) interface{}

	// Visit a parse tree produced by ApexParser#booleanLiteralValue.
	VisitBooleanLiteralValue(ctx *BooleanLiteralValueContext) interface{}

	// Visit a parse tree produced by ApexParser#signedNumberValue.
	VisitSignedNumberValue(ctx *SignedNumberValueContext) interface{}

	// Visit a parse tree produced by ApexParser#stringLiteralValue.
	VisitStringLiteralValue(ctx *StringLiteralValueContext) interface{}

	// Visit a parse tree produced by ApexParser#dateLiteralValue.
	VisitDateLiteralValue(ctx *DateLiteralValueContext) interface{}

	// Visit a parse tree produced by ApexParser#dateTimeLiteralValue.
	VisitDateTimeLiteralValue(ctx *DateTimeLiteralValueContext) interface{}

	// Visit a parse tree produced by ApexParser#dateFormulaValue.
	VisitDateFormulaValue(ctx *DateFormulaValueContext) interface{}

	// Visit a parse tree produced by ApexParser#currencyValueValue.
	VisitCurrencyValueValue(ctx *CurrencyValueValueContext) interface{}

	// Visit a parse tree produced by ApexParser#subQueryValue.
	VisitSubQueryValue(ctx *SubQueryValueContext) interface{}

	// Visit a parse tree produced by ApexParser#valueListValue.
	VisitValueListValue(ctx *ValueListValueContext) interface{}

	// Visit a parse tree produced by ApexParser#boundExpressionValue.
	VisitBoundExpressionValue(ctx *BoundExpressionValueContext) interface{}

	// Visit a parse tree produced by ApexParser#valueList.
	VisitValueList(ctx *ValueListContext) interface{}

	// Visit a parse tree produced by ApexParser#currencyValue.
	VisitCurrencyValue(ctx *CurrencyValueContext) interface{}

	// Visit a parse tree produced by ApexParser#signedNumber.
	VisitSignedNumber(ctx *SignedNumberContext) interface{}

	// Visit a parse tree produced by ApexParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by ApexParser#filteringExpression.
	VisitFilteringExpression(ctx *FilteringExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#dataCategorySelection.
	VisitDataCategorySelection(ctx *DataCategorySelectionContext) interface{}

	// Visit a parse tree produced by ApexParser#dataCategoryName.
	VisitDataCategoryName(ctx *DataCategoryNameContext) interface{}

	// Visit a parse tree produced by ApexParser#filteringSelector.
	VisitFilteringSelector(ctx *FilteringSelectorContext) interface{}

	// Visit a parse tree produced by ApexParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by ApexParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by ApexParser#fieldOrderList.
	VisitFieldOrderList(ctx *FieldOrderListContext) interface{}

	// Visit a parse tree produced by ApexParser#fieldOrder.
	VisitFieldOrder(ctx *FieldOrderContext) interface{}

	// Visit a parse tree produced by ApexParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by ApexParser#offsetClause.
	VisitOffsetClause(ctx *OffsetClauseContext) interface{}

	// Visit a parse tree produced by ApexParser#allRowsClause.
	VisitAllRowsClause(ctx *AllRowsClauseContext) interface{}

	// Visit a parse tree produced by ApexParser#forClauses.
	VisitForClauses(ctx *ForClausesContext) interface{}

	// Visit a parse tree produced by ApexParser#forClause.
	VisitForClause(ctx *ForClauseContext) interface{}

	// Visit a parse tree produced by ApexParser#boundExpression.
	VisitBoundExpression(ctx *BoundExpressionContext) interface{}

	// Visit a parse tree produced by ApexParser#dateFormula.
	VisitDateFormula(ctx *DateFormulaContext) interface{}

	// Visit a parse tree produced by ApexParser#signedInteger.
	VisitSignedInteger(ctx *SignedIntegerContext) interface{}

	// Visit a parse tree produced by ApexParser#soqlId.
	VisitSoqlId(ctx *SoqlIdContext) interface{}

	// Visit a parse tree produced by ApexParser#soslLiteral.
	VisitSoslLiteral(ctx *SoslLiteralContext) interface{}

	// Visit a parse tree produced by ApexParser#soslLiteralAlt.
	VisitSoslLiteralAlt(ctx *SoslLiteralAltContext) interface{}

	// Visit a parse tree produced by ApexParser#soslClauses.
	VisitSoslClauses(ctx *SoslClausesContext) interface{}

	// Visit a parse tree produced by ApexParser#inSearchGroup.
	VisitInSearchGroup(ctx *InSearchGroupContext) interface{}

	// Visit a parse tree produced by ApexParser#returningFieldSpecList.
	VisitReturningFieldSpecList(ctx *ReturningFieldSpecListContext) interface{}

	// Visit a parse tree produced by ApexParser#withDivisionAssign.
	VisitWithDivisionAssign(ctx *WithDivisionAssignContext) interface{}

	// Visit a parse tree produced by ApexParser#withDataCategory.
	VisitWithDataCategory(ctx *WithDataCategoryContext) interface{}

	// Visit a parse tree produced by ApexParser#withSnippet.
	VisitWithSnippet(ctx *WithSnippetContext) interface{}

	// Visit a parse tree produced by ApexParser#withNetworkIn.
	VisitWithNetworkIn(ctx *WithNetworkInContext) interface{}

	// Visit a parse tree produced by ApexParser#withNetworkAssign.
	VisitWithNetworkAssign(ctx *WithNetworkAssignContext) interface{}

	// Visit a parse tree produced by ApexParser#withPricebookIdAssign.
	VisitWithPricebookIdAssign(ctx *WithPricebookIdAssignContext) interface{}

	// Visit a parse tree produced by ApexParser#withMetadataAssign.
	VisitWithMetadataAssign(ctx *WithMetadataAssignContext) interface{}

	// Visit a parse tree produced by ApexParser#updateListClause.
	VisitUpdateListClause(ctx *UpdateListClauseContext) interface{}

	// Visit a parse tree produced by ApexParser#searchGroup.
	VisitSearchGroup(ctx *SearchGroupContext) interface{}

	// Visit a parse tree produced by ApexParser#fieldSpecList.
	VisitFieldSpecList(ctx *FieldSpecListContext) interface{}

	// Visit a parse tree produced by ApexParser#fieldSpec.
	VisitFieldSpec(ctx *FieldSpecContext) interface{}

	// Visit a parse tree produced by ApexParser#fieldSpecClauses.
	VisitFieldSpecClauses(ctx *FieldSpecClausesContext) interface{}

	// Visit a parse tree produced by ApexParser#fieldList.
	VisitFieldList(ctx *FieldListContext) interface{}

	// Visit a parse tree produced by ApexParser#updateList.
	VisitUpdateList(ctx *UpdateListContext) interface{}

	// Visit a parse tree produced by ApexParser#updateType.
	VisitUpdateType(ctx *UpdateTypeContext) interface{}

	// Visit a parse tree produced by ApexParser#networkList.
	VisitNetworkList(ctx *NetworkListContext) interface{}

	// Visit a parse tree produced by ApexParser#soslId.
	VisitSoslId(ctx *SoslIdContext) interface{}

	// Visit a parse tree produced by ApexParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by ApexParser#anyId.
	VisitAnyId(ctx *AnyIdContext) interface{}
}
