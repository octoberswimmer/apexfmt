// Code generated from ./ApexParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ApexParser
import "github.com/antlr4-go/antlr/v4"

// BaseApexParserListener is a complete listener for a parse tree produced by ApexParser.
type BaseApexParserListener struct{}

var _ ApexParserListener = &BaseApexParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseApexParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseApexParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseApexParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseApexParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTriggerUnit is called when production triggerUnit is entered.
func (s *BaseApexParserListener) EnterTriggerUnit(ctx *TriggerUnitContext) {}

// ExitTriggerUnit is called when production triggerUnit is exited.
func (s *BaseApexParserListener) ExitTriggerUnit(ctx *TriggerUnitContext) {}

// EnterTriggerCase is called when production triggerCase is entered.
func (s *BaseApexParserListener) EnterTriggerCase(ctx *TriggerCaseContext) {}

// ExitTriggerCase is called when production triggerCase is exited.
func (s *BaseApexParserListener) ExitTriggerCase(ctx *TriggerCaseContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseApexParserListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseApexParserListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseApexParserListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseApexParserListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseApexParserListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseApexParserListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterEnumDeclaration is called when production enumDeclaration is entered.
func (s *BaseApexParserListener) EnterEnumDeclaration(ctx *EnumDeclarationContext) {}

// ExitEnumDeclaration is called when production enumDeclaration is exited.
func (s *BaseApexParserListener) ExitEnumDeclaration(ctx *EnumDeclarationContext) {}

// EnterEnumConstants is called when production enumConstants is entered.
func (s *BaseApexParserListener) EnterEnumConstants(ctx *EnumConstantsContext) {}

// ExitEnumConstants is called when production enumConstants is exited.
func (s *BaseApexParserListener) ExitEnumConstants(ctx *EnumConstantsContext) {}

// EnterInterfaceDeclaration is called when production interfaceDeclaration is entered.
func (s *BaseApexParserListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// ExitInterfaceDeclaration is called when production interfaceDeclaration is exited.
func (s *BaseApexParserListener) ExitInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseApexParserListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseApexParserListener) ExitTypeList(ctx *TypeListContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseApexParserListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseApexParserListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterInterfaceBody is called when production interfaceBody is entered.
func (s *BaseApexParserListener) EnterInterfaceBody(ctx *InterfaceBodyContext) {}

// ExitInterfaceBody is called when production interfaceBody is exited.
func (s *BaseApexParserListener) ExitInterfaceBody(ctx *InterfaceBodyContext) {}

// EnterClassBodyDeclaration is called when production classBodyDeclaration is entered.
func (s *BaseApexParserListener) EnterClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// ExitClassBodyDeclaration is called when production classBodyDeclaration is exited.
func (s *BaseApexParserListener) ExitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseApexParserListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseApexParserListener) ExitModifier(ctx *ModifierContext) {}

// EnterMemberDeclaration is called when production memberDeclaration is entered.
func (s *BaseApexParserListener) EnterMemberDeclaration(ctx *MemberDeclarationContext) {}

// ExitMemberDeclaration is called when production memberDeclaration is exited.
func (s *BaseApexParserListener) ExitMemberDeclaration(ctx *MemberDeclarationContext) {}

// EnterMethodDeclaration is called when production methodDeclaration is entered.
func (s *BaseApexParserListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {}

// ExitMethodDeclaration is called when production methodDeclaration is exited.
func (s *BaseApexParserListener) ExitMethodDeclaration(ctx *MethodDeclarationContext) {}

// EnterConstructorDeclaration is called when production constructorDeclaration is entered.
func (s *BaseApexParserListener) EnterConstructorDeclaration(ctx *ConstructorDeclarationContext) {}

// ExitConstructorDeclaration is called when production constructorDeclaration is exited.
func (s *BaseApexParserListener) ExitConstructorDeclaration(ctx *ConstructorDeclarationContext) {}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BaseApexParserListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BaseApexParserListener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterPropertyDeclaration is called when production propertyDeclaration is entered.
func (s *BaseApexParserListener) EnterPropertyDeclaration(ctx *PropertyDeclarationContext) {}

// ExitPropertyDeclaration is called when production propertyDeclaration is exited.
func (s *BaseApexParserListener) ExitPropertyDeclaration(ctx *PropertyDeclarationContext) {}

// EnterInterfaceMethodDeclaration is called when production interfaceMethodDeclaration is entered.
func (s *BaseApexParserListener) EnterInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {
}

// ExitInterfaceMethodDeclaration is called when production interfaceMethodDeclaration is exited.
func (s *BaseApexParserListener) ExitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {
}

// EnterVariableDeclarators is called when production variableDeclarators is entered.
func (s *BaseApexParserListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// ExitVariableDeclarators is called when production variableDeclarators is exited.
func (s *BaseApexParserListener) ExitVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// EnterVariableDeclarator is called when production variableDeclarator is entered.
func (s *BaseApexParserListener) EnterVariableDeclarator(ctx *VariableDeclaratorContext) {}

// ExitVariableDeclarator is called when production variableDeclarator is exited.
func (s *BaseApexParserListener) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {}

// EnterArrayInitializer is called when production arrayInitializer is entered.
func (s *BaseApexParserListener) EnterArrayInitializer(ctx *ArrayInitializerContext) {}

// ExitArrayInitializer is called when production arrayInitializer is exited.
func (s *BaseApexParserListener) ExitArrayInitializer(ctx *ArrayInitializerContext) {}

// EnterTypeRef is called when production typeRef is entered.
func (s *BaseApexParserListener) EnterTypeRef(ctx *TypeRefContext) {}

// ExitTypeRef is called when production typeRef is exited.
func (s *BaseApexParserListener) ExitTypeRef(ctx *TypeRefContext) {}

// EnterArraySubscripts is called when production arraySubscripts is entered.
func (s *BaseApexParserListener) EnterArraySubscripts(ctx *ArraySubscriptsContext) {}

// ExitArraySubscripts is called when production arraySubscripts is exited.
func (s *BaseApexParserListener) ExitArraySubscripts(ctx *ArraySubscriptsContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseApexParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseApexParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterTypeArguments is called when production typeArguments is entered.
func (s *BaseApexParserListener) EnterTypeArguments(ctx *TypeArgumentsContext) {}

// ExitTypeArguments is called when production typeArguments is exited.
func (s *BaseApexParserListener) ExitTypeArguments(ctx *TypeArgumentsContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BaseApexParserListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BaseApexParserListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseApexParserListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseApexParserListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BaseApexParserListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BaseApexParserListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseApexParserListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseApexParserListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseApexParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseApexParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseApexParserListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseApexParserListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterElementValuePairs is called when production elementValuePairs is entered.
func (s *BaseApexParserListener) EnterElementValuePairs(ctx *ElementValuePairsContext) {}

// ExitElementValuePairs is called when production elementValuePairs is exited.
func (s *BaseApexParserListener) ExitElementValuePairs(ctx *ElementValuePairsContext) {}

// EnterDelimitedElementValuePair is called when production delimitedElementValuePair is entered.
func (s *BaseApexParserListener) EnterDelimitedElementValuePair(ctx *DelimitedElementValuePairContext) {
}

// ExitDelimitedElementValuePair is called when production delimitedElementValuePair is exited.
func (s *BaseApexParserListener) ExitDelimitedElementValuePair(ctx *DelimitedElementValuePairContext) {
}

// EnterElementValuePair is called when production elementValuePair is entered.
func (s *BaseApexParserListener) EnterElementValuePair(ctx *ElementValuePairContext) {}

// ExitElementValuePair is called when production elementValuePair is exited.
func (s *BaseApexParserListener) ExitElementValuePair(ctx *ElementValuePairContext) {}

// EnterElementValue is called when production elementValue is entered.
func (s *BaseApexParserListener) EnterElementValue(ctx *ElementValueContext) {}

// ExitElementValue is called when production elementValue is exited.
func (s *BaseApexParserListener) ExitElementValue(ctx *ElementValueContext) {}

// EnterElementValueArrayInitializer is called when production elementValueArrayInitializer is entered.
func (s *BaseApexParserListener) EnterElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) {
}

// ExitElementValueArrayInitializer is called when production elementValueArrayInitializer is exited.
func (s *BaseApexParserListener) ExitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) {
}

// EnterTrailingComma is called when production trailingComma is entered.
func (s *BaseApexParserListener) EnterTrailingComma(ctx *TrailingCommaContext) {}

// ExitTrailingComma is called when production trailingComma is exited.
func (s *BaseApexParserListener) ExitTrailingComma(ctx *TrailingCommaContext) {}

// EnterTriggerBlock is called when production triggerBlock is entered.
func (s *BaseApexParserListener) EnterTriggerBlock(ctx *TriggerBlockContext) {}

// ExitTriggerBlock is called when production triggerBlock is exited.
func (s *BaseApexParserListener) ExitTriggerBlock(ctx *TriggerBlockContext) {}

// EnterTriggerStatement is called when production triggerStatement is entered.
func (s *BaseApexParserListener) EnterTriggerStatement(ctx *TriggerStatementContext) {}

// ExitTriggerStatement is called when production triggerStatement is exited.
func (s *BaseApexParserListener) ExitTriggerStatement(ctx *TriggerStatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseApexParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseApexParserListener) ExitBlock(ctx *BlockContext) {}

// EnterLocalVariableDeclarationStatement is called when production localVariableDeclarationStatement is entered.
func (s *BaseApexParserListener) EnterLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) {
}

// ExitLocalVariableDeclarationStatement is called when production localVariableDeclarationStatement is exited.
func (s *BaseApexParserListener) ExitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) {
}

// EnterLocalVariableDeclaration is called when production localVariableDeclaration is entered.
func (s *BaseApexParserListener) EnterLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {
}

// ExitLocalVariableDeclaration is called when production localVariableDeclaration is exited.
func (s *BaseApexParserListener) ExitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseApexParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseApexParserListener) ExitStatement(ctx *StatementContext) {}

// EnterBlockMemberDeclaration is called when production blockMemberDeclaration is entered.
func (s *BaseApexParserListener) EnterBlockMemberDeclaration(ctx *BlockMemberDeclarationContext) {}

// ExitBlockMemberDeclaration is called when production blockMemberDeclaration is exited.
func (s *BaseApexParserListener) ExitBlockMemberDeclaration(ctx *BlockMemberDeclarationContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseApexParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseApexParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseApexParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseApexParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterWhenControl is called when production whenControl is entered.
func (s *BaseApexParserListener) EnterWhenControl(ctx *WhenControlContext) {}

// ExitWhenControl is called when production whenControl is exited.
func (s *BaseApexParserListener) ExitWhenControl(ctx *WhenControlContext) {}

// EnterWhenValue is called when production whenValue is entered.
func (s *BaseApexParserListener) EnterWhenValue(ctx *WhenValueContext) {}

// ExitWhenValue is called when production whenValue is exited.
func (s *BaseApexParserListener) ExitWhenValue(ctx *WhenValueContext) {}

// EnterWhenLiteral is called when production whenLiteral is entered.
func (s *BaseApexParserListener) EnterWhenLiteral(ctx *WhenLiteralContext) {}

// ExitWhenLiteral is called when production whenLiteral is exited.
func (s *BaseApexParserListener) ExitWhenLiteral(ctx *WhenLiteralContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseApexParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseApexParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseApexParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseApexParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterDoWhileStatement is called when production doWhileStatement is entered.
func (s *BaseApexParserListener) EnterDoWhileStatement(ctx *DoWhileStatementContext) {}

// ExitDoWhileStatement is called when production doWhileStatement is exited.
func (s *BaseApexParserListener) ExitDoWhileStatement(ctx *DoWhileStatementContext) {}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseApexParserListener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseApexParserListener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseApexParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseApexParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterThrowStatement is called when production throwStatement is entered.
func (s *BaseApexParserListener) EnterThrowStatement(ctx *ThrowStatementContext) {}

// ExitThrowStatement is called when production throwStatement is exited.
func (s *BaseApexParserListener) ExitThrowStatement(ctx *ThrowStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseApexParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseApexParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseApexParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseApexParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterInsertStatement is called when production insertStatement is entered.
func (s *BaseApexParserListener) EnterInsertStatement(ctx *InsertStatementContext) {}

// ExitInsertStatement is called when production insertStatement is exited.
func (s *BaseApexParserListener) ExitInsertStatement(ctx *InsertStatementContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseApexParserListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseApexParserListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseApexParserListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseApexParserListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterUndeleteStatement is called when production undeleteStatement is entered.
func (s *BaseApexParserListener) EnterUndeleteStatement(ctx *UndeleteStatementContext) {}

// ExitUndeleteStatement is called when production undeleteStatement is exited.
func (s *BaseApexParserListener) ExitUndeleteStatement(ctx *UndeleteStatementContext) {}

// EnterUpsertStatement is called when production upsertStatement is entered.
func (s *BaseApexParserListener) EnterUpsertStatement(ctx *UpsertStatementContext) {}

// ExitUpsertStatement is called when production upsertStatement is exited.
func (s *BaseApexParserListener) ExitUpsertStatement(ctx *UpsertStatementContext) {}

// EnterMergeStatement is called when production mergeStatement is entered.
func (s *BaseApexParserListener) EnterMergeStatement(ctx *MergeStatementContext) {}

// ExitMergeStatement is called when production mergeStatement is exited.
func (s *BaseApexParserListener) ExitMergeStatement(ctx *MergeStatementContext) {}

// EnterRunAsStatement is called when production runAsStatement is entered.
func (s *BaseApexParserListener) EnterRunAsStatement(ctx *RunAsStatementContext) {}

// ExitRunAsStatement is called when production runAsStatement is exited.
func (s *BaseApexParserListener) ExitRunAsStatement(ctx *RunAsStatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseApexParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseApexParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterPropertyBlock is called when production propertyBlock is entered.
func (s *BaseApexParserListener) EnterPropertyBlock(ctx *PropertyBlockContext) {}

// ExitPropertyBlock is called when production propertyBlock is exited.
func (s *BaseApexParserListener) ExitPropertyBlock(ctx *PropertyBlockContext) {}

// EnterGetter is called when production getter is entered.
func (s *BaseApexParserListener) EnterGetter(ctx *GetterContext) {}

// ExitGetter is called when production getter is exited.
func (s *BaseApexParserListener) ExitGetter(ctx *GetterContext) {}

// EnterSetter is called when production setter is entered.
func (s *BaseApexParserListener) EnterSetter(ctx *SetterContext) {}

// ExitSetter is called when production setter is exited.
func (s *BaseApexParserListener) ExitSetter(ctx *SetterContext) {}

// EnterCatchClause is called when production catchClause is entered.
func (s *BaseApexParserListener) EnterCatchClause(ctx *CatchClauseContext) {}

// ExitCatchClause is called when production catchClause is exited.
func (s *BaseApexParserListener) ExitCatchClause(ctx *CatchClauseContext) {}

// EnterFinallyBlock is called when production finallyBlock is entered.
func (s *BaseApexParserListener) EnterFinallyBlock(ctx *FinallyBlockContext) {}

// ExitFinallyBlock is called when production finallyBlock is exited.
func (s *BaseApexParserListener) ExitFinallyBlock(ctx *FinallyBlockContext) {}

// EnterForControl is called when production forControl is entered.
func (s *BaseApexParserListener) EnterForControl(ctx *ForControlContext) {}

// ExitForControl is called when production forControl is exited.
func (s *BaseApexParserListener) ExitForControl(ctx *ForControlContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BaseApexParserListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseApexParserListener) ExitForInit(ctx *ForInitContext) {}

// EnterEnhancedForControl is called when production enhancedForControl is entered.
func (s *BaseApexParserListener) EnterEnhancedForControl(ctx *EnhancedForControlContext) {}

// ExitEnhancedForControl is called when production enhancedForControl is exited.
func (s *BaseApexParserListener) ExitEnhancedForControl(ctx *EnhancedForControlContext) {}

// EnterForUpdate is called when production forUpdate is entered.
func (s *BaseApexParserListener) EnterForUpdate(ctx *ForUpdateContext) {}

// ExitForUpdate is called when production forUpdate is exited.
func (s *BaseApexParserListener) ExitForUpdate(ctx *ForUpdateContext) {}

// EnterParExpression is called when production parExpression is entered.
func (s *BaseApexParserListener) EnterParExpression(ctx *ParExpressionContext) {}

// ExitParExpression is called when production parExpression is exited.
func (s *BaseApexParserListener) ExitParExpression(ctx *ParExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseApexParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseApexParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseApexParserListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseApexParserListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterArth1Expression is called when production arth1Expression is entered.
func (s *BaseApexParserListener) EnterArth1Expression(ctx *Arth1ExpressionContext) {}

// ExitArth1Expression is called when production arth1Expression is exited.
func (s *BaseApexParserListener) ExitArth1Expression(ctx *Arth1ExpressionContext) {}

// EnterCoalExpression is called when production coalExpression is entered.
func (s *BaseApexParserListener) EnterCoalExpression(ctx *CoalExpressionContext) {}

// ExitCoalExpression is called when production coalExpression is exited.
func (s *BaseApexParserListener) ExitCoalExpression(ctx *CoalExpressionContext) {}

// EnterDotExpression is called when production dotExpression is entered.
func (s *BaseApexParserListener) EnterDotExpression(ctx *DotExpressionContext) {}

// ExitDotExpression is called when production dotExpression is exited.
func (s *BaseApexParserListener) ExitDotExpression(ctx *DotExpressionContext) {}

// EnterBitOrExpression is called when production bitOrExpression is entered.
func (s *BaseApexParserListener) EnterBitOrExpression(ctx *BitOrExpressionContext) {}

// ExitBitOrExpression is called when production bitOrExpression is exited.
func (s *BaseApexParserListener) ExitBitOrExpression(ctx *BitOrExpressionContext) {}

// EnterArrayExpression is called when production arrayExpression is entered.
func (s *BaseApexParserListener) EnterArrayExpression(ctx *ArrayExpressionContext) {}

// ExitArrayExpression is called when production arrayExpression is exited.
func (s *BaseApexParserListener) ExitArrayExpression(ctx *ArrayExpressionContext) {}

// EnterAssignExpression is called when production assignExpression is entered.
func (s *BaseApexParserListener) EnterAssignExpression(ctx *AssignExpressionContext) {}

// ExitAssignExpression is called when production assignExpression is exited.
func (s *BaseApexParserListener) ExitAssignExpression(ctx *AssignExpressionContext) {}

// EnterMethodCallExpression is called when production methodCallExpression is entered.
func (s *BaseApexParserListener) EnterMethodCallExpression(ctx *MethodCallExpressionContext) {}

// ExitMethodCallExpression is called when production methodCallExpression is exited.
func (s *BaseApexParserListener) ExitMethodCallExpression(ctx *MethodCallExpressionContext) {}

// EnterBitNotExpression is called when production bitNotExpression is entered.
func (s *BaseApexParserListener) EnterBitNotExpression(ctx *BitNotExpressionContext) {}

// ExitBitNotExpression is called when production bitNotExpression is exited.
func (s *BaseApexParserListener) ExitBitNotExpression(ctx *BitNotExpressionContext) {}

// EnterNewInstanceExpression is called when production newInstanceExpression is entered.
func (s *BaseApexParserListener) EnterNewInstanceExpression(ctx *NewInstanceExpressionContext) {}

// ExitNewInstanceExpression is called when production newInstanceExpression is exited.
func (s *BaseApexParserListener) ExitNewInstanceExpression(ctx *NewInstanceExpressionContext) {}

// EnterArth2Expression is called when production arth2Expression is entered.
func (s *BaseApexParserListener) EnterArth2Expression(ctx *Arth2ExpressionContext) {}

// ExitArth2Expression is called when production arth2Expression is exited.
func (s *BaseApexParserListener) ExitArth2Expression(ctx *Arth2ExpressionContext) {}

// EnterLogAndExpression is called when production logAndExpression is entered.
func (s *BaseApexParserListener) EnterLogAndExpression(ctx *LogAndExpressionContext) {}

// ExitLogAndExpression is called when production logAndExpression is exited.
func (s *BaseApexParserListener) ExitLogAndExpression(ctx *LogAndExpressionContext) {}

// EnterCastExpression is called when production castExpression is entered.
func (s *BaseApexParserListener) EnterCastExpression(ctx *CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *BaseApexParserListener) ExitCastExpression(ctx *CastExpressionContext) {}

// EnterBitAndExpression is called when production bitAndExpression is entered.
func (s *BaseApexParserListener) EnterBitAndExpression(ctx *BitAndExpressionContext) {}

// ExitBitAndExpression is called when production bitAndExpression is exited.
func (s *BaseApexParserListener) ExitBitAndExpression(ctx *BitAndExpressionContext) {}

// EnterCmpExpression is called when production cmpExpression is entered.
func (s *BaseApexParserListener) EnterCmpExpression(ctx *CmpExpressionContext) {}

// ExitCmpExpression is called when production cmpExpression is exited.
func (s *BaseApexParserListener) ExitCmpExpression(ctx *CmpExpressionContext) {}

// EnterBitExpression is called when production bitExpression is entered.
func (s *BaseApexParserListener) EnterBitExpression(ctx *BitExpressionContext) {}

// ExitBitExpression is called when production bitExpression is exited.
func (s *BaseApexParserListener) ExitBitExpression(ctx *BitExpressionContext) {}

// EnterLogOrExpression is called when production logOrExpression is entered.
func (s *BaseApexParserListener) EnterLogOrExpression(ctx *LogOrExpressionContext) {}

// ExitLogOrExpression is called when production logOrExpression is exited.
func (s *BaseApexParserListener) ExitLogOrExpression(ctx *LogOrExpressionContext) {}

// EnterCondExpression is called when production condExpression is entered.
func (s *BaseApexParserListener) EnterCondExpression(ctx *CondExpressionContext) {}

// ExitCondExpression is called when production condExpression is exited.
func (s *BaseApexParserListener) ExitCondExpression(ctx *CondExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseApexParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseApexParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterPostOpExpression is called when production postOpExpression is entered.
func (s *BaseApexParserListener) EnterPostOpExpression(ctx *PostOpExpressionContext) {}

// ExitPostOpExpression is called when production postOpExpression is exited.
func (s *BaseApexParserListener) ExitPostOpExpression(ctx *PostOpExpressionContext) {}

// EnterNegExpression is called when production negExpression is entered.
func (s *BaseApexParserListener) EnterNegExpression(ctx *NegExpressionContext) {}

// ExitNegExpression is called when production negExpression is exited.
func (s *BaseApexParserListener) ExitNegExpression(ctx *NegExpressionContext) {}

// EnterPreOpExpression is called when production preOpExpression is entered.
func (s *BaseApexParserListener) EnterPreOpExpression(ctx *PreOpExpressionContext) {}

// ExitPreOpExpression is called when production preOpExpression is exited.
func (s *BaseApexParserListener) ExitPreOpExpression(ctx *PreOpExpressionContext) {}

// EnterSubExpression is called when production subExpression is entered.
func (s *BaseApexParserListener) EnterSubExpression(ctx *SubExpressionContext) {}

// ExitSubExpression is called when production subExpression is exited.
func (s *BaseApexParserListener) ExitSubExpression(ctx *SubExpressionContext) {}

// EnterInstanceOfExpression is called when production instanceOfExpression is entered.
func (s *BaseApexParserListener) EnterInstanceOfExpression(ctx *InstanceOfExpressionContext) {}

// ExitInstanceOfExpression is called when production instanceOfExpression is exited.
func (s *BaseApexParserListener) ExitInstanceOfExpression(ctx *InstanceOfExpressionContext) {}

// EnterThisPrimary is called when production thisPrimary is entered.
func (s *BaseApexParserListener) EnterThisPrimary(ctx *ThisPrimaryContext) {}

// ExitThisPrimary is called when production thisPrimary is exited.
func (s *BaseApexParserListener) ExitThisPrimary(ctx *ThisPrimaryContext) {}

// EnterSuperPrimary is called when production superPrimary is entered.
func (s *BaseApexParserListener) EnterSuperPrimary(ctx *SuperPrimaryContext) {}

// ExitSuperPrimary is called when production superPrimary is exited.
func (s *BaseApexParserListener) ExitSuperPrimary(ctx *SuperPrimaryContext) {}

// EnterLiteralPrimary is called when production literalPrimary is entered.
func (s *BaseApexParserListener) EnterLiteralPrimary(ctx *LiteralPrimaryContext) {}

// ExitLiteralPrimary is called when production literalPrimary is exited.
func (s *BaseApexParserListener) ExitLiteralPrimary(ctx *LiteralPrimaryContext) {}

// EnterTypeRefPrimary is called when production typeRefPrimary is entered.
func (s *BaseApexParserListener) EnterTypeRefPrimary(ctx *TypeRefPrimaryContext) {}

// ExitTypeRefPrimary is called when production typeRefPrimary is exited.
func (s *BaseApexParserListener) ExitTypeRefPrimary(ctx *TypeRefPrimaryContext) {}

// EnterIdPrimary is called when production idPrimary is entered.
func (s *BaseApexParserListener) EnterIdPrimary(ctx *IdPrimaryContext) {}

// ExitIdPrimary is called when production idPrimary is exited.
func (s *BaseApexParserListener) ExitIdPrimary(ctx *IdPrimaryContext) {}

// EnterSoqlPrimary is called when production soqlPrimary is entered.
func (s *BaseApexParserListener) EnterSoqlPrimary(ctx *SoqlPrimaryContext) {}

// ExitSoqlPrimary is called when production soqlPrimary is exited.
func (s *BaseApexParserListener) ExitSoqlPrimary(ctx *SoqlPrimaryContext) {}

// EnterSoslPrimary is called when production soslPrimary is entered.
func (s *BaseApexParserListener) EnterSoslPrimary(ctx *SoslPrimaryContext) {}

// ExitSoslPrimary is called when production soslPrimary is exited.
func (s *BaseApexParserListener) ExitSoslPrimary(ctx *SoslPrimaryContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BaseApexParserListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BaseApexParserListener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterDotMethodCall is called when production dotMethodCall is entered.
func (s *BaseApexParserListener) EnterDotMethodCall(ctx *DotMethodCallContext) {}

// ExitDotMethodCall is called when production dotMethodCall is exited.
func (s *BaseApexParserListener) ExitDotMethodCall(ctx *DotMethodCallContext) {}

// EnterCreator is called when production creator is entered.
func (s *BaseApexParserListener) EnterCreator(ctx *CreatorContext) {}

// ExitCreator is called when production creator is exited.
func (s *BaseApexParserListener) ExitCreator(ctx *CreatorContext) {}

// EnterCreatedName is called when production createdName is entered.
func (s *BaseApexParserListener) EnterCreatedName(ctx *CreatedNameContext) {}

// ExitCreatedName is called when production createdName is exited.
func (s *BaseApexParserListener) ExitCreatedName(ctx *CreatedNameContext) {}

// EnterIdCreatedNamePair is called when production idCreatedNamePair is entered.
func (s *BaseApexParserListener) EnterIdCreatedNamePair(ctx *IdCreatedNamePairContext) {}

// ExitIdCreatedNamePair is called when production idCreatedNamePair is exited.
func (s *BaseApexParserListener) ExitIdCreatedNamePair(ctx *IdCreatedNamePairContext) {}

// EnterNoRest is called when production noRest is entered.
func (s *BaseApexParserListener) EnterNoRest(ctx *NoRestContext) {}

// ExitNoRest is called when production noRest is exited.
func (s *BaseApexParserListener) ExitNoRest(ctx *NoRestContext) {}

// EnterClassCreatorRest is called when production classCreatorRest is entered.
func (s *BaseApexParserListener) EnterClassCreatorRest(ctx *ClassCreatorRestContext) {}

// ExitClassCreatorRest is called when production classCreatorRest is exited.
func (s *BaseApexParserListener) ExitClassCreatorRest(ctx *ClassCreatorRestContext) {}

// EnterArrayCreatorRest is called when production arrayCreatorRest is entered.
func (s *BaseApexParserListener) EnterArrayCreatorRest(ctx *ArrayCreatorRestContext) {}

// ExitArrayCreatorRest is called when production arrayCreatorRest is exited.
func (s *BaseApexParserListener) ExitArrayCreatorRest(ctx *ArrayCreatorRestContext) {}

// EnterMapCreatorRest is called when production mapCreatorRest is entered.
func (s *BaseApexParserListener) EnterMapCreatorRest(ctx *MapCreatorRestContext) {}

// ExitMapCreatorRest is called when production mapCreatorRest is exited.
func (s *BaseApexParserListener) ExitMapCreatorRest(ctx *MapCreatorRestContext) {}

// EnterMapCreatorRestPair is called when production mapCreatorRestPair is entered.
func (s *BaseApexParserListener) EnterMapCreatorRestPair(ctx *MapCreatorRestPairContext) {}

// ExitMapCreatorRestPair is called when production mapCreatorRestPair is exited.
func (s *BaseApexParserListener) ExitMapCreatorRestPair(ctx *MapCreatorRestPairContext) {}

// EnterSetCreatorRest is called when production setCreatorRest is entered.
func (s *BaseApexParserListener) EnterSetCreatorRest(ctx *SetCreatorRestContext) {}

// ExitSetCreatorRest is called when production setCreatorRest is exited.
func (s *BaseApexParserListener) ExitSetCreatorRest(ctx *SetCreatorRestContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseApexParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseApexParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterSoqlLiteral is called when production soqlLiteral is entered.
func (s *BaseApexParserListener) EnterSoqlLiteral(ctx *SoqlLiteralContext) {}

// ExitSoqlLiteral is called when production soqlLiteral is exited.
func (s *BaseApexParserListener) ExitSoqlLiteral(ctx *SoqlLiteralContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseApexParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseApexParserListener) ExitQuery(ctx *QueryContext) {}

// EnterSubQuery is called when production subQuery is entered.
func (s *BaseApexParserListener) EnterSubQuery(ctx *SubQueryContext) {}

// ExitSubQuery is called when production subQuery is exited.
func (s *BaseApexParserListener) ExitSubQuery(ctx *SubQueryContext) {}

// EnterSelectList is called when production selectList is entered.
func (s *BaseApexParserListener) EnterSelectList(ctx *SelectListContext) {}

// ExitSelectList is called when production selectList is exited.
func (s *BaseApexParserListener) ExitSelectList(ctx *SelectListContext) {}

// EnterSelectEntry is called when production selectEntry is entered.
func (s *BaseApexParserListener) EnterSelectEntry(ctx *SelectEntryContext) {}

// ExitSelectEntry is called when production selectEntry is exited.
func (s *BaseApexParserListener) ExitSelectEntry(ctx *SelectEntryContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseApexParserListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseApexParserListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterFromNameList is called when production fromNameList is entered.
func (s *BaseApexParserListener) EnterFromNameList(ctx *FromNameListContext) {}

// ExitFromNameList is called when production fromNameList is exited.
func (s *BaseApexParserListener) ExitFromNameList(ctx *FromNameListContext) {}

// EnterFieldNameAlias is called when production fieldNameAlias is entered.
func (s *BaseApexParserListener) EnterFieldNameAlias(ctx *FieldNameAliasContext) {}

// ExitFieldNameAlias is called when production fieldNameAlias is exited.
func (s *BaseApexParserListener) ExitFieldNameAlias(ctx *FieldNameAliasContext) {}

// EnterSubFieldList is called when production subFieldList is entered.
func (s *BaseApexParserListener) EnterSubFieldList(ctx *SubFieldListContext) {}

// ExitSubFieldList is called when production subFieldList is exited.
func (s *BaseApexParserListener) ExitSubFieldList(ctx *SubFieldListContext) {}

// EnterSubFieldEntry is called when production subFieldEntry is entered.
func (s *BaseApexParserListener) EnterSubFieldEntry(ctx *SubFieldEntryContext) {}

// ExitSubFieldEntry is called when production subFieldEntry is exited.
func (s *BaseApexParserListener) ExitSubFieldEntry(ctx *SubFieldEntryContext) {}

// EnterSoqlFieldsParameter is called when production soqlFieldsParameter is entered.
func (s *BaseApexParserListener) EnterSoqlFieldsParameter(ctx *SoqlFieldsParameterContext) {}

// ExitSoqlFieldsParameter is called when production soqlFieldsParameter is exited.
func (s *BaseApexParserListener) ExitSoqlFieldsParameter(ctx *SoqlFieldsParameterContext) {}

// EnterSoqlFunction is called when production soqlFunction is entered.
func (s *BaseApexParserListener) EnterSoqlFunction(ctx *SoqlFunctionContext) {}

// ExitSoqlFunction is called when production soqlFunction is exited.
func (s *BaseApexParserListener) ExitSoqlFunction(ctx *SoqlFunctionContext) {}

// EnterDateFieldName is called when production dateFieldName is entered.
func (s *BaseApexParserListener) EnterDateFieldName(ctx *DateFieldNameContext) {}

// ExitDateFieldName is called when production dateFieldName is exited.
func (s *BaseApexParserListener) ExitDateFieldName(ctx *DateFieldNameContext) {}

// EnterLocationValue is called when production locationValue is entered.
func (s *BaseApexParserListener) EnterLocationValue(ctx *LocationValueContext) {}

// ExitLocationValue is called when production locationValue is exited.
func (s *BaseApexParserListener) ExitLocationValue(ctx *LocationValueContext) {}

// EnterCoordinateValue is called when production coordinateValue is entered.
func (s *BaseApexParserListener) EnterCoordinateValue(ctx *CoordinateValueContext) {}

// ExitCoordinateValue is called when production coordinateValue is exited.
func (s *BaseApexParserListener) ExitCoordinateValue(ctx *CoordinateValueContext) {}

// EnterTypeOf is called when production typeOf is entered.
func (s *BaseApexParserListener) EnterTypeOf(ctx *TypeOfContext) {}

// ExitTypeOf is called when production typeOf is exited.
func (s *BaseApexParserListener) ExitTypeOf(ctx *TypeOfContext) {}

// EnterWhenClause is called when production whenClause is entered.
func (s *BaseApexParserListener) EnterWhenClause(ctx *WhenClauseContext) {}

// ExitWhenClause is called when production whenClause is exited.
func (s *BaseApexParserListener) ExitWhenClause(ctx *WhenClauseContext) {}

// EnterElseClause is called when production elseClause is entered.
func (s *BaseApexParserListener) EnterElseClause(ctx *ElseClauseContext) {}

// ExitElseClause is called when production elseClause is exited.
func (s *BaseApexParserListener) ExitElseClause(ctx *ElseClauseContext) {}

// EnterFieldNameList is called when production fieldNameList is entered.
func (s *BaseApexParserListener) EnterFieldNameList(ctx *FieldNameListContext) {}

// ExitFieldNameList is called when production fieldNameList is exited.
func (s *BaseApexParserListener) ExitFieldNameList(ctx *FieldNameListContext) {}

// EnterUsingScope is called when production usingScope is entered.
func (s *BaseApexParserListener) EnterUsingScope(ctx *UsingScopeContext) {}

// ExitUsingScope is called when production usingScope is exited.
func (s *BaseApexParserListener) ExitUsingScope(ctx *UsingScopeContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseApexParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseApexParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterLogicalExpression is called when production logicalExpression is entered.
func (s *BaseApexParserListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production logicalExpression is exited.
func (s *BaseApexParserListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BaseApexParserListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BaseApexParserListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterFieldExpression is called when production fieldExpression is entered.
func (s *BaseApexParserListener) EnterFieldExpression(ctx *FieldExpressionContext) {}

// ExitFieldExpression is called when production fieldExpression is exited.
func (s *BaseApexParserListener) ExitFieldExpression(ctx *FieldExpressionContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseApexParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseApexParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterNullValue is called when production nullValue is entered.
func (s *BaseApexParserListener) EnterNullValue(ctx *NullValueContext) {}

// ExitNullValue is called when production nullValue is exited.
func (s *BaseApexParserListener) ExitNullValue(ctx *NullValueContext) {}

// EnterBooleanLiteralValue is called when production booleanLiteralValue is entered.
func (s *BaseApexParserListener) EnterBooleanLiteralValue(ctx *BooleanLiteralValueContext) {}

// ExitBooleanLiteralValue is called when production booleanLiteralValue is exited.
func (s *BaseApexParserListener) ExitBooleanLiteralValue(ctx *BooleanLiteralValueContext) {}

// EnterSignedNumberValue is called when production signedNumberValue is entered.
func (s *BaseApexParserListener) EnterSignedNumberValue(ctx *SignedNumberValueContext) {}

// ExitSignedNumberValue is called when production signedNumberValue is exited.
func (s *BaseApexParserListener) ExitSignedNumberValue(ctx *SignedNumberValueContext) {}

// EnterStringLiteralValue is called when production stringLiteralValue is entered.
func (s *BaseApexParserListener) EnterStringLiteralValue(ctx *StringLiteralValueContext) {}

// ExitStringLiteralValue is called when production stringLiteralValue is exited.
func (s *BaseApexParserListener) ExitStringLiteralValue(ctx *StringLiteralValueContext) {}

// EnterDateLiteralValue is called when production dateLiteralValue is entered.
func (s *BaseApexParserListener) EnterDateLiteralValue(ctx *DateLiteralValueContext) {}

// ExitDateLiteralValue is called when production dateLiteralValue is exited.
func (s *BaseApexParserListener) ExitDateLiteralValue(ctx *DateLiteralValueContext) {}

// EnterTimeLiteralValue is called when production timeLiteralValue is entered.
func (s *BaseApexParserListener) EnterTimeLiteralValue(ctx *TimeLiteralValueContext) {}

// ExitTimeLiteralValue is called when production timeLiteralValue is exited.
func (s *BaseApexParserListener) ExitTimeLiteralValue(ctx *TimeLiteralValueContext) {}

// EnterDateTimeLiteralValue is called when production dateTimeLiteralValue is entered.
func (s *BaseApexParserListener) EnterDateTimeLiteralValue(ctx *DateTimeLiteralValueContext) {}

// ExitDateTimeLiteralValue is called when production dateTimeLiteralValue is exited.
func (s *BaseApexParserListener) ExitDateTimeLiteralValue(ctx *DateTimeLiteralValueContext) {}

// EnterDateFormulaValue is called when production dateFormulaValue is entered.
func (s *BaseApexParserListener) EnterDateFormulaValue(ctx *DateFormulaValueContext) {}

// ExitDateFormulaValue is called when production dateFormulaValue is exited.
func (s *BaseApexParserListener) ExitDateFormulaValue(ctx *DateFormulaValueContext) {}

// EnterCurrencyValueValue is called when production currencyValueValue is entered.
func (s *BaseApexParserListener) EnterCurrencyValueValue(ctx *CurrencyValueValueContext) {}

// ExitCurrencyValueValue is called when production currencyValueValue is exited.
func (s *BaseApexParserListener) ExitCurrencyValueValue(ctx *CurrencyValueValueContext) {}

// EnterSubQueryValue is called when production subQueryValue is entered.
func (s *BaseApexParserListener) EnterSubQueryValue(ctx *SubQueryValueContext) {}

// ExitSubQueryValue is called when production subQueryValue is exited.
func (s *BaseApexParserListener) ExitSubQueryValue(ctx *SubQueryValueContext) {}

// EnterValueListValue is called when production valueListValue is entered.
func (s *BaseApexParserListener) EnterValueListValue(ctx *ValueListValueContext) {}

// ExitValueListValue is called when production valueListValue is exited.
func (s *BaseApexParserListener) ExitValueListValue(ctx *ValueListValueContext) {}

// EnterBoundExpressionValue is called when production boundExpressionValue is entered.
func (s *BaseApexParserListener) EnterBoundExpressionValue(ctx *BoundExpressionValueContext) {}

// ExitBoundExpressionValue is called when production boundExpressionValue is exited.
func (s *BaseApexParserListener) ExitBoundExpressionValue(ctx *BoundExpressionValueContext) {}

// EnterValueList is called when production valueList is entered.
func (s *BaseApexParserListener) EnterValueList(ctx *ValueListContext) {}

// ExitValueList is called when production valueList is exited.
func (s *BaseApexParserListener) ExitValueList(ctx *ValueListContext) {}

// EnterCurrencyValue is called when production currencyValue is entered.
func (s *BaseApexParserListener) EnterCurrencyValue(ctx *CurrencyValueContext) {}

// ExitCurrencyValue is called when production currencyValue is exited.
func (s *BaseApexParserListener) ExitCurrencyValue(ctx *CurrencyValueContext) {}

// EnterSignedNumber is called when production signedNumber is entered.
func (s *BaseApexParserListener) EnterSignedNumber(ctx *SignedNumberContext) {}

// ExitSignedNumber is called when production signedNumber is exited.
func (s *BaseApexParserListener) ExitSignedNumber(ctx *SignedNumberContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseApexParserListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseApexParserListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterFilteringExpression is called when production filteringExpression is entered.
func (s *BaseApexParserListener) EnterFilteringExpression(ctx *FilteringExpressionContext) {}

// ExitFilteringExpression is called when production filteringExpression is exited.
func (s *BaseApexParserListener) ExitFilteringExpression(ctx *FilteringExpressionContext) {}

// EnterDataCategorySelection is called when production dataCategorySelection is entered.
func (s *BaseApexParserListener) EnterDataCategorySelection(ctx *DataCategorySelectionContext) {}

// ExitDataCategorySelection is called when production dataCategorySelection is exited.
func (s *BaseApexParserListener) ExitDataCategorySelection(ctx *DataCategorySelectionContext) {}

// EnterDataCategoryName is called when production dataCategoryName is entered.
func (s *BaseApexParserListener) EnterDataCategoryName(ctx *DataCategoryNameContext) {}

// ExitDataCategoryName is called when production dataCategoryName is exited.
func (s *BaseApexParserListener) ExitDataCategoryName(ctx *DataCategoryNameContext) {}

// EnterFilteringSelector is called when production filteringSelector is entered.
func (s *BaseApexParserListener) EnterFilteringSelector(ctx *FilteringSelectorContext) {}

// ExitFilteringSelector is called when production filteringSelector is exited.
func (s *BaseApexParserListener) ExitFilteringSelector(ctx *FilteringSelectorContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseApexParserListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseApexParserListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseApexParserListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseApexParserListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterFieldOrderList is called when production fieldOrderList is entered.
func (s *BaseApexParserListener) EnterFieldOrderList(ctx *FieldOrderListContext) {}

// ExitFieldOrderList is called when production fieldOrderList is exited.
func (s *BaseApexParserListener) ExitFieldOrderList(ctx *FieldOrderListContext) {}

// EnterFieldOrder is called when production fieldOrder is entered.
func (s *BaseApexParserListener) EnterFieldOrder(ctx *FieldOrderContext) {}

// ExitFieldOrder is called when production fieldOrder is exited.
func (s *BaseApexParserListener) ExitFieldOrder(ctx *FieldOrderContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseApexParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseApexParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterOffsetClause is called when production offsetClause is entered.
func (s *BaseApexParserListener) EnterOffsetClause(ctx *OffsetClauseContext) {}

// ExitOffsetClause is called when production offsetClause is exited.
func (s *BaseApexParserListener) ExitOffsetClause(ctx *OffsetClauseContext) {}

// EnterAllRowsClause is called when production allRowsClause is entered.
func (s *BaseApexParserListener) EnterAllRowsClause(ctx *AllRowsClauseContext) {}

// ExitAllRowsClause is called when production allRowsClause is exited.
func (s *BaseApexParserListener) ExitAllRowsClause(ctx *AllRowsClauseContext) {}

// EnterForClauses is called when production forClauses is entered.
func (s *BaseApexParserListener) EnterForClauses(ctx *ForClausesContext) {}

// ExitForClauses is called when production forClauses is exited.
func (s *BaseApexParserListener) ExitForClauses(ctx *ForClausesContext) {}

// EnterForClause is called when production forClause is entered.
func (s *BaseApexParserListener) EnterForClause(ctx *ForClauseContext) {}

// ExitForClause is called when production forClause is exited.
func (s *BaseApexParserListener) ExitForClause(ctx *ForClauseContext) {}

// EnterBoundExpression is called when production boundExpression is entered.
func (s *BaseApexParserListener) EnterBoundExpression(ctx *BoundExpressionContext) {}

// ExitBoundExpression is called when production boundExpression is exited.
func (s *BaseApexParserListener) ExitBoundExpression(ctx *BoundExpressionContext) {}

// EnterDateFormula is called when production dateFormula is entered.
func (s *BaseApexParserListener) EnterDateFormula(ctx *DateFormulaContext) {}

// ExitDateFormula is called when production dateFormula is exited.
func (s *BaseApexParserListener) ExitDateFormula(ctx *DateFormulaContext) {}

// EnterSignedInteger is called when production signedInteger is entered.
func (s *BaseApexParserListener) EnterSignedInteger(ctx *SignedIntegerContext) {}

// ExitSignedInteger is called when production signedInteger is exited.
func (s *BaseApexParserListener) ExitSignedInteger(ctx *SignedIntegerContext) {}

// EnterSoqlId is called when production soqlId is entered.
func (s *BaseApexParserListener) EnterSoqlId(ctx *SoqlIdContext) {}

// ExitSoqlId is called when production soqlId is exited.
func (s *BaseApexParserListener) ExitSoqlId(ctx *SoqlIdContext) {}

// EnterSoslLiteral is called when production soslLiteral is entered.
func (s *BaseApexParserListener) EnterSoslLiteral(ctx *SoslLiteralContext) {}

// ExitSoslLiteral is called when production soslLiteral is exited.
func (s *BaseApexParserListener) ExitSoslLiteral(ctx *SoslLiteralContext) {}

// EnterSoslLiteralAlt is called when production soslLiteralAlt is entered.
func (s *BaseApexParserListener) EnterSoslLiteralAlt(ctx *SoslLiteralAltContext) {}

// ExitSoslLiteralAlt is called when production soslLiteralAlt is exited.
func (s *BaseApexParserListener) ExitSoslLiteralAlt(ctx *SoslLiteralAltContext) {}

// EnterSoslClauses is called when production soslClauses is entered.
func (s *BaseApexParserListener) EnterSoslClauses(ctx *SoslClausesContext) {}

// ExitSoslClauses is called when production soslClauses is exited.
func (s *BaseApexParserListener) ExitSoslClauses(ctx *SoslClausesContext) {}

// EnterInSearchGroup is called when production inSearchGroup is entered.
func (s *BaseApexParserListener) EnterInSearchGroup(ctx *InSearchGroupContext) {}

// ExitInSearchGroup is called when production inSearchGroup is exited.
func (s *BaseApexParserListener) ExitInSearchGroup(ctx *InSearchGroupContext) {}

// EnterReturningFieldSpecList is called when production returningFieldSpecList is entered.
func (s *BaseApexParserListener) EnterReturningFieldSpecList(ctx *ReturningFieldSpecListContext) {}

// ExitReturningFieldSpecList is called when production returningFieldSpecList is exited.
func (s *BaseApexParserListener) ExitReturningFieldSpecList(ctx *ReturningFieldSpecListContext) {}

// EnterWithDivisionAssign is called when production withDivisionAssign is entered.
func (s *BaseApexParserListener) EnterWithDivisionAssign(ctx *WithDivisionAssignContext) {}

// ExitWithDivisionAssign is called when production withDivisionAssign is exited.
func (s *BaseApexParserListener) ExitWithDivisionAssign(ctx *WithDivisionAssignContext) {}

// EnterWithDataCategory is called when production withDataCategory is entered.
func (s *BaseApexParserListener) EnterWithDataCategory(ctx *WithDataCategoryContext) {}

// ExitWithDataCategory is called when production withDataCategory is exited.
func (s *BaseApexParserListener) ExitWithDataCategory(ctx *WithDataCategoryContext) {}

// EnterWithSnippet is called when production withSnippet is entered.
func (s *BaseApexParserListener) EnterWithSnippet(ctx *WithSnippetContext) {}

// ExitWithSnippet is called when production withSnippet is exited.
func (s *BaseApexParserListener) ExitWithSnippet(ctx *WithSnippetContext) {}

// EnterWithNetworkIn is called when production withNetworkIn is entered.
func (s *BaseApexParserListener) EnterWithNetworkIn(ctx *WithNetworkInContext) {}

// ExitWithNetworkIn is called when production withNetworkIn is exited.
func (s *BaseApexParserListener) ExitWithNetworkIn(ctx *WithNetworkInContext) {}

// EnterWithNetworkAssign is called when production withNetworkAssign is entered.
func (s *BaseApexParserListener) EnterWithNetworkAssign(ctx *WithNetworkAssignContext) {}

// ExitWithNetworkAssign is called when production withNetworkAssign is exited.
func (s *BaseApexParserListener) ExitWithNetworkAssign(ctx *WithNetworkAssignContext) {}

// EnterWithPricebookIdAssign is called when production withPricebookIdAssign is entered.
func (s *BaseApexParserListener) EnterWithPricebookIdAssign(ctx *WithPricebookIdAssignContext) {}

// ExitWithPricebookIdAssign is called when production withPricebookIdAssign is exited.
func (s *BaseApexParserListener) ExitWithPricebookIdAssign(ctx *WithPricebookIdAssignContext) {}

// EnterWithMetadataAssign is called when production withMetadataAssign is entered.
func (s *BaseApexParserListener) EnterWithMetadataAssign(ctx *WithMetadataAssignContext) {}

// ExitWithMetadataAssign is called when production withMetadataAssign is exited.
func (s *BaseApexParserListener) ExitWithMetadataAssign(ctx *WithMetadataAssignContext) {}

// EnterWithModeClause is called when production withModeClause is entered.
func (s *BaseApexParserListener) EnterWithModeClause(ctx *WithModeClauseContext) {}

// ExitWithModeClause is called when production withModeClause is exited.
func (s *BaseApexParserListener) ExitWithModeClause(ctx *WithModeClauseContext) {}

// EnterUpdateListClause is called when production updateListClause is entered.
func (s *BaseApexParserListener) EnterUpdateListClause(ctx *UpdateListClauseContext) {}

// ExitUpdateListClause is called when production updateListClause is exited.
func (s *BaseApexParserListener) ExitUpdateListClause(ctx *UpdateListClauseContext) {}

// EnterSearchGroup is called when production searchGroup is entered.
func (s *BaseApexParserListener) EnterSearchGroup(ctx *SearchGroupContext) {}

// ExitSearchGroup is called when production searchGroup is exited.
func (s *BaseApexParserListener) ExitSearchGroup(ctx *SearchGroupContext) {}

// EnterFieldSpecList is called when production fieldSpecList is entered.
func (s *BaseApexParserListener) EnterFieldSpecList(ctx *FieldSpecListContext) {}

// ExitFieldSpecList is called when production fieldSpecList is exited.
func (s *BaseApexParserListener) ExitFieldSpecList(ctx *FieldSpecListContext) {}

// EnterFieldSpec is called when production fieldSpec is entered.
func (s *BaseApexParserListener) EnterFieldSpec(ctx *FieldSpecContext) {}

// ExitFieldSpec is called when production fieldSpec is exited.
func (s *BaseApexParserListener) ExitFieldSpec(ctx *FieldSpecContext) {}

// EnterFieldSpecClauses is called when production fieldSpecClauses is entered.
func (s *BaseApexParserListener) EnterFieldSpecClauses(ctx *FieldSpecClausesContext) {}

// ExitFieldSpecClauses is called when production fieldSpecClauses is exited.
func (s *BaseApexParserListener) ExitFieldSpecClauses(ctx *FieldSpecClausesContext) {}

// EnterFieldList is called when production fieldList is entered.
func (s *BaseApexParserListener) EnterFieldList(ctx *FieldListContext) {}

// ExitFieldList is called when production fieldList is exited.
func (s *BaseApexParserListener) ExitFieldList(ctx *FieldListContext) {}

// EnterUpdateList is called when production updateList is entered.
func (s *BaseApexParserListener) EnterUpdateList(ctx *UpdateListContext) {}

// ExitUpdateList is called when production updateList is exited.
func (s *BaseApexParserListener) ExitUpdateList(ctx *UpdateListContext) {}

// EnterUpdateType is called when production updateType is entered.
func (s *BaseApexParserListener) EnterUpdateType(ctx *UpdateTypeContext) {}

// ExitUpdateType is called when production updateType is exited.
func (s *BaseApexParserListener) ExitUpdateType(ctx *UpdateTypeContext) {}

// EnterNetworkList is called when production networkList is entered.
func (s *BaseApexParserListener) EnterNetworkList(ctx *NetworkListContext) {}

// ExitNetworkList is called when production networkList is exited.
func (s *BaseApexParserListener) ExitNetworkList(ctx *NetworkListContext) {}

// EnterSoslId is called when production soslId is entered.
func (s *BaseApexParserListener) EnterSoslId(ctx *SoslIdContext) {}

// ExitSoslId is called when production soslId is exited.
func (s *BaseApexParserListener) ExitSoslId(ctx *SoslIdContext) {}

// EnterId is called when production id is entered.
func (s *BaseApexParserListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseApexParserListener) ExitId(ctx *IdContext) {}

// EnterMethodId is called when production methodId is entered.
func (s *BaseApexParserListener) EnterMethodId(ctx *MethodIdContext) {}

// ExitMethodId is called when production methodId is exited.
func (s *BaseApexParserListener) ExitMethodId(ctx *MethodIdContext) {}

// EnterAnyId is called when production anyId is entered.
func (s *BaseApexParserListener) EnterAnyId(ctx *AnyIdContext) {}

// ExitAnyId is called when production anyId is exited.
func (s *BaseApexParserListener) ExitAnyId(ctx *AnyIdContext) {}
