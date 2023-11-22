// Code generated from ./apex.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // apex
import "github.com/antlr4-go/antlr/v4"

// BaseapexListener is a complete listener for a parse tree produced by apexParser.
type BaseapexListener struct{}

var _ apexListener = &BaseapexListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseapexListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseapexListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseapexListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseapexListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseapexListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseapexListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterPackageDeclaration is called when production packageDeclaration is entered.
func (s *BaseapexListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {}

// ExitPackageDeclaration is called when production packageDeclaration is exited.
func (s *BaseapexListener) ExitPackageDeclaration(ctx *PackageDeclarationContext) {}

// EnterImportDeclaration is called when production importDeclaration is entered.
func (s *BaseapexListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseapexListener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseapexListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseapexListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseapexListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseapexListener) ExitModifier(ctx *ModifierContext) {}

// EnterClassOrInterfaceModifier is called when production classOrInterfaceModifier is entered.
func (s *BaseapexListener) EnterClassOrInterfaceModifier(ctx *ClassOrInterfaceModifierContext) {}

// ExitClassOrInterfaceModifier is called when production classOrInterfaceModifier is exited.
func (s *BaseapexListener) ExitClassOrInterfaceModifier(ctx *ClassOrInterfaceModifierContext) {}

// EnterVariableModifier is called when production variableModifier is entered.
func (s *BaseapexListener) EnterVariableModifier(ctx *VariableModifierContext) {}

// ExitVariableModifier is called when production variableModifier is exited.
func (s *BaseapexListener) ExitVariableModifier(ctx *VariableModifierContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseapexListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseapexListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterTypeParameters is called when production typeParameters is entered.
func (s *BaseapexListener) EnterTypeParameters(ctx *TypeParametersContext) {}

// ExitTypeParameters is called when production typeParameters is exited.
func (s *BaseapexListener) ExitTypeParameters(ctx *TypeParametersContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseapexListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseapexListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterTypeBound is called when production typeBound is entered.
func (s *BaseapexListener) EnterTypeBound(ctx *TypeBoundContext) {}

// ExitTypeBound is called when production typeBound is exited.
func (s *BaseapexListener) ExitTypeBound(ctx *TypeBoundContext) {}

// EnterEnumDeclaration is called when production enumDeclaration is entered.
func (s *BaseapexListener) EnterEnumDeclaration(ctx *EnumDeclarationContext) {}

// ExitEnumDeclaration is called when production enumDeclaration is exited.
func (s *BaseapexListener) ExitEnumDeclaration(ctx *EnumDeclarationContext) {}

// EnterEnumConstants is called when production enumConstants is entered.
func (s *BaseapexListener) EnterEnumConstants(ctx *EnumConstantsContext) {}

// ExitEnumConstants is called when production enumConstants is exited.
func (s *BaseapexListener) ExitEnumConstants(ctx *EnumConstantsContext) {}

// EnterEnumConstant is called when production enumConstant is entered.
func (s *BaseapexListener) EnterEnumConstant(ctx *EnumConstantContext) {}

// ExitEnumConstant is called when production enumConstant is exited.
func (s *BaseapexListener) ExitEnumConstant(ctx *EnumConstantContext) {}

// EnterEnumBodyDeclarations is called when production enumBodyDeclarations is entered.
func (s *BaseapexListener) EnterEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) {}

// ExitEnumBodyDeclarations is called when production enumBodyDeclarations is exited.
func (s *BaseapexListener) ExitEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) {}

// EnterInterfaceDeclaration is called when production interfaceDeclaration is entered.
func (s *BaseapexListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// ExitInterfaceDeclaration is called when production interfaceDeclaration is exited.
func (s *BaseapexListener) ExitInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseapexListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseapexListener) ExitTypeList(ctx *TypeListContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseapexListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseapexListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterInterfaceBody is called when production interfaceBody is entered.
func (s *BaseapexListener) EnterInterfaceBody(ctx *InterfaceBodyContext) {}

// ExitInterfaceBody is called when production interfaceBody is exited.
func (s *BaseapexListener) ExitInterfaceBody(ctx *InterfaceBodyContext) {}

// EnterClassBodyDeclaration is called when production classBodyDeclaration is entered.
func (s *BaseapexListener) EnterClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// ExitClassBodyDeclaration is called when production classBodyDeclaration is exited.
func (s *BaseapexListener) ExitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// EnterMemberDeclaration is called when production memberDeclaration is entered.
func (s *BaseapexListener) EnterMemberDeclaration(ctx *MemberDeclarationContext) {}

// ExitMemberDeclaration is called when production memberDeclaration is exited.
func (s *BaseapexListener) ExitMemberDeclaration(ctx *MemberDeclarationContext) {}

// EnterMethodDeclaration is called when production methodDeclaration is entered.
func (s *BaseapexListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {}

// ExitMethodDeclaration is called when production methodDeclaration is exited.
func (s *BaseapexListener) ExitMethodDeclaration(ctx *MethodDeclarationContext) {}

// EnterGenericMethodDeclaration is called when production genericMethodDeclaration is entered.
func (s *BaseapexListener) EnterGenericMethodDeclaration(ctx *GenericMethodDeclarationContext) {}

// ExitGenericMethodDeclaration is called when production genericMethodDeclaration is exited.
func (s *BaseapexListener) ExitGenericMethodDeclaration(ctx *GenericMethodDeclarationContext) {}

// EnterConstructorDeclaration is called when production constructorDeclaration is entered.
func (s *BaseapexListener) EnterConstructorDeclaration(ctx *ConstructorDeclarationContext) {}

// ExitConstructorDeclaration is called when production constructorDeclaration is exited.
func (s *BaseapexListener) ExitConstructorDeclaration(ctx *ConstructorDeclarationContext) {}

// EnterGenericConstructorDeclaration is called when production genericConstructorDeclaration is entered.
func (s *BaseapexListener) EnterGenericConstructorDeclaration(ctx *GenericConstructorDeclarationContext) {
}

// ExitGenericConstructorDeclaration is called when production genericConstructorDeclaration is exited.
func (s *BaseapexListener) ExitGenericConstructorDeclaration(ctx *GenericConstructorDeclarationContext) {
}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BaseapexListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BaseapexListener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterPropertyDeclaration is called when production propertyDeclaration is entered.
func (s *BaseapexListener) EnterPropertyDeclaration(ctx *PropertyDeclarationContext) {}

// ExitPropertyDeclaration is called when production propertyDeclaration is exited.
func (s *BaseapexListener) ExitPropertyDeclaration(ctx *PropertyDeclarationContext) {}

// EnterPropertyBodyDeclaration is called when production propertyBodyDeclaration is entered.
func (s *BaseapexListener) EnterPropertyBodyDeclaration(ctx *PropertyBodyDeclarationContext) {}

// ExitPropertyBodyDeclaration is called when production propertyBodyDeclaration is exited.
func (s *BaseapexListener) ExitPropertyBodyDeclaration(ctx *PropertyBodyDeclarationContext) {}

// EnterInterfaceBodyDeclaration is called when production interfaceBodyDeclaration is entered.
func (s *BaseapexListener) EnterInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) {}

// ExitInterfaceBodyDeclaration is called when production interfaceBodyDeclaration is exited.
func (s *BaseapexListener) ExitInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) {}

// EnterInterfaceMemberDeclaration is called when production interfaceMemberDeclaration is entered.
func (s *BaseapexListener) EnterInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) {}

// ExitInterfaceMemberDeclaration is called when production interfaceMemberDeclaration is exited.
func (s *BaseapexListener) ExitInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) {}

// EnterConstDeclaration is called when production constDeclaration is entered.
func (s *BaseapexListener) EnterConstDeclaration(ctx *ConstDeclarationContext) {}

// ExitConstDeclaration is called when production constDeclaration is exited.
func (s *BaseapexListener) ExitConstDeclaration(ctx *ConstDeclarationContext) {}

// EnterConstantDeclarator is called when production constantDeclarator is entered.
func (s *BaseapexListener) EnterConstantDeclarator(ctx *ConstantDeclaratorContext) {}

// ExitConstantDeclarator is called when production constantDeclarator is exited.
func (s *BaseapexListener) ExitConstantDeclarator(ctx *ConstantDeclaratorContext) {}

// EnterInterfaceMethodDeclaration is called when production interfaceMethodDeclaration is entered.
func (s *BaseapexListener) EnterInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {}

// ExitInterfaceMethodDeclaration is called when production interfaceMethodDeclaration is exited.
func (s *BaseapexListener) ExitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {}

// EnterGenericInterfaceMethodDeclaration is called when production genericInterfaceMethodDeclaration is entered.
func (s *BaseapexListener) EnterGenericInterfaceMethodDeclaration(ctx *GenericInterfaceMethodDeclarationContext) {
}

// ExitGenericInterfaceMethodDeclaration is called when production genericInterfaceMethodDeclaration is exited.
func (s *BaseapexListener) ExitGenericInterfaceMethodDeclaration(ctx *GenericInterfaceMethodDeclarationContext) {
}

// EnterVariableDeclarators is called when production variableDeclarators is entered.
func (s *BaseapexListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// ExitVariableDeclarators is called when production variableDeclarators is exited.
func (s *BaseapexListener) ExitVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// EnterVariableDeclarator is called when production variableDeclarator is entered.
func (s *BaseapexListener) EnterVariableDeclarator(ctx *VariableDeclaratorContext) {}

// ExitVariableDeclarator is called when production variableDeclarator is exited.
func (s *BaseapexListener) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {}

// EnterVariableDeclaratorId is called when production variableDeclaratorId is entered.
func (s *BaseapexListener) EnterVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// ExitVariableDeclaratorId is called when production variableDeclaratorId is exited.
func (s *BaseapexListener) ExitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BaseapexListener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BaseapexListener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

// EnterArrayInitializer is called when production arrayInitializer is entered.
func (s *BaseapexListener) EnterArrayInitializer(ctx *ArrayInitializerContext) {}

// ExitArrayInitializer is called when production arrayInitializer is exited.
func (s *BaseapexListener) ExitArrayInitializer(ctx *ArrayInitializerContext) {}

// EnterEnumConstantName is called when production enumConstantName is entered.
func (s *BaseapexListener) EnterEnumConstantName(ctx *EnumConstantNameContext) {}

// ExitEnumConstantName is called when production enumConstantName is exited.
func (s *BaseapexListener) ExitEnumConstantName(ctx *EnumConstantNameContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseapexListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseapexListener) ExitType_(ctx *Type_Context) {}

// EnterClassOrInterfaceType is called when production classOrInterfaceType is entered.
func (s *BaseapexListener) EnterClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// ExitClassOrInterfaceType is called when production classOrInterfaceType is exited.
func (s *BaseapexListener) ExitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseapexListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseapexListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterTypeArguments is called when production typeArguments is entered.
func (s *BaseapexListener) EnterTypeArguments(ctx *TypeArgumentsContext) {}

// ExitTypeArguments is called when production typeArguments is exited.
func (s *BaseapexListener) ExitTypeArguments(ctx *TypeArgumentsContext) {}

// EnterTypeArgument is called when production typeArgument is entered.
func (s *BaseapexListener) EnterTypeArgument(ctx *TypeArgumentContext) {}

// ExitTypeArgument is called when production typeArgument is exited.
func (s *BaseapexListener) ExitTypeArgument(ctx *TypeArgumentContext) {}

// EnterQualifiedNameList is called when production qualifiedNameList is entered.
func (s *BaseapexListener) EnterQualifiedNameList(ctx *QualifiedNameListContext) {}

// ExitQualifiedNameList is called when production qualifiedNameList is exited.
func (s *BaseapexListener) ExitQualifiedNameList(ctx *QualifiedNameListContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BaseapexListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BaseapexListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseapexListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseapexListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BaseapexListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BaseapexListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterLastFormalParameter is called when production lastFormalParameter is entered.
func (s *BaseapexListener) EnterLastFormalParameter(ctx *LastFormalParameterContext) {}

// ExitLastFormalParameter is called when production lastFormalParameter is exited.
func (s *BaseapexListener) ExitLastFormalParameter(ctx *LastFormalParameterContext) {}

// EnterMethodBody is called when production methodBody is entered.
func (s *BaseapexListener) EnterMethodBody(ctx *MethodBodyContext) {}

// ExitMethodBody is called when production methodBody is exited.
func (s *BaseapexListener) ExitMethodBody(ctx *MethodBodyContext) {}

// EnterConstructorBody is called when production constructorBody is entered.
func (s *BaseapexListener) EnterConstructorBody(ctx *ConstructorBodyContext) {}

// ExitConstructorBody is called when production constructorBody is exited.
func (s *BaseapexListener) ExitConstructorBody(ctx *ConstructorBodyContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseapexListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseapexListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseapexListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseapexListener) ExitLiteral(ctx *LiteralContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseapexListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseapexListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterAnnotationName is called when production annotationName is entered.
func (s *BaseapexListener) EnterAnnotationName(ctx *AnnotationNameContext) {}

// ExitAnnotationName is called when production annotationName is exited.
func (s *BaseapexListener) ExitAnnotationName(ctx *AnnotationNameContext) {}

// EnterElementValuePairs is called when production elementValuePairs is entered.
func (s *BaseapexListener) EnterElementValuePairs(ctx *ElementValuePairsContext) {}

// ExitElementValuePairs is called when production elementValuePairs is exited.
func (s *BaseapexListener) ExitElementValuePairs(ctx *ElementValuePairsContext) {}

// EnterElementValuePair is called when production elementValuePair is entered.
func (s *BaseapexListener) EnterElementValuePair(ctx *ElementValuePairContext) {}

// ExitElementValuePair is called when production elementValuePair is exited.
func (s *BaseapexListener) ExitElementValuePair(ctx *ElementValuePairContext) {}

// EnterElementValue is called when production elementValue is entered.
func (s *BaseapexListener) EnterElementValue(ctx *ElementValueContext) {}

// ExitElementValue is called when production elementValue is exited.
func (s *BaseapexListener) ExitElementValue(ctx *ElementValueContext) {}

// EnterElementValueArrayInitializer is called when production elementValueArrayInitializer is entered.
func (s *BaseapexListener) EnterElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) {
}

// ExitElementValueArrayInitializer is called when production elementValueArrayInitializer is exited.
func (s *BaseapexListener) ExitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) {
}

// EnterAnnotationTypeDeclaration is called when production annotationTypeDeclaration is entered.
func (s *BaseapexListener) EnterAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) {}

// ExitAnnotationTypeDeclaration is called when production annotationTypeDeclaration is exited.
func (s *BaseapexListener) ExitAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) {}

// EnterAnnotationTypeBody is called when production annotationTypeBody is entered.
func (s *BaseapexListener) EnterAnnotationTypeBody(ctx *AnnotationTypeBodyContext) {}

// ExitAnnotationTypeBody is called when production annotationTypeBody is exited.
func (s *BaseapexListener) ExitAnnotationTypeBody(ctx *AnnotationTypeBodyContext) {}

// EnterAnnotationTypeElementDeclaration is called when production annotationTypeElementDeclaration is entered.
func (s *BaseapexListener) EnterAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) {
}

// ExitAnnotationTypeElementDeclaration is called when production annotationTypeElementDeclaration is exited.
func (s *BaseapexListener) ExitAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) {
}

// EnterAnnotationTypeElementRest is called when production annotationTypeElementRest is entered.
func (s *BaseapexListener) EnterAnnotationTypeElementRest(ctx *AnnotationTypeElementRestContext) {}

// ExitAnnotationTypeElementRest is called when production annotationTypeElementRest is exited.
func (s *BaseapexListener) ExitAnnotationTypeElementRest(ctx *AnnotationTypeElementRestContext) {}

// EnterAnnotationMethodOrConstantRest is called when production annotationMethodOrConstantRest is entered.
func (s *BaseapexListener) EnterAnnotationMethodOrConstantRest(ctx *AnnotationMethodOrConstantRestContext) {
}

// ExitAnnotationMethodOrConstantRest is called when production annotationMethodOrConstantRest is exited.
func (s *BaseapexListener) ExitAnnotationMethodOrConstantRest(ctx *AnnotationMethodOrConstantRestContext) {
}

// EnterAnnotationMethodRest is called when production annotationMethodRest is entered.
func (s *BaseapexListener) EnterAnnotationMethodRest(ctx *AnnotationMethodRestContext) {}

// ExitAnnotationMethodRest is called when production annotationMethodRest is exited.
func (s *BaseapexListener) ExitAnnotationMethodRest(ctx *AnnotationMethodRestContext) {}

// EnterAnnotationConstantRest is called when production annotationConstantRest is entered.
func (s *BaseapexListener) EnterAnnotationConstantRest(ctx *AnnotationConstantRestContext) {}

// ExitAnnotationConstantRest is called when production annotationConstantRest is exited.
func (s *BaseapexListener) ExitAnnotationConstantRest(ctx *AnnotationConstantRestContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseapexListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseapexListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseapexListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseapexListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseapexListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseapexListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterLocalVariableDeclarationStatement is called when production localVariableDeclarationStatement is entered.
func (s *BaseapexListener) EnterLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) {
}

// ExitLocalVariableDeclarationStatement is called when production localVariableDeclarationStatement is exited.
func (s *BaseapexListener) ExitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) {
}

// EnterLocalVariableDeclaration is called when production localVariableDeclaration is entered.
func (s *BaseapexListener) EnterLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {}

// ExitLocalVariableDeclaration is called when production localVariableDeclaration is exited.
func (s *BaseapexListener) ExitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseapexListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseapexListener) ExitStatement(ctx *StatementContext) {}

// EnterPropertyBlock is called when production propertyBlock is entered.
func (s *BaseapexListener) EnterPropertyBlock(ctx *PropertyBlockContext) {}

// ExitPropertyBlock is called when production propertyBlock is exited.
func (s *BaseapexListener) ExitPropertyBlock(ctx *PropertyBlockContext) {}

// EnterGetter is called when production getter is entered.
func (s *BaseapexListener) EnterGetter(ctx *GetterContext) {}

// ExitGetter is called when production getter is exited.
func (s *BaseapexListener) ExitGetter(ctx *GetterContext) {}

// EnterSetter is called when production setter is entered.
func (s *BaseapexListener) EnterSetter(ctx *SetterContext) {}

// ExitSetter is called when production setter is exited.
func (s *BaseapexListener) ExitSetter(ctx *SetterContext) {}

// EnterCatchClause is called when production catchClause is entered.
func (s *BaseapexListener) EnterCatchClause(ctx *CatchClauseContext) {}

// ExitCatchClause is called when production catchClause is exited.
func (s *BaseapexListener) ExitCatchClause(ctx *CatchClauseContext) {}

// EnterCatchType is called when production catchType is entered.
func (s *BaseapexListener) EnterCatchType(ctx *CatchTypeContext) {}

// ExitCatchType is called when production catchType is exited.
func (s *BaseapexListener) ExitCatchType(ctx *CatchTypeContext) {}

// EnterFinallyBlock is called when production finallyBlock is entered.
func (s *BaseapexListener) EnterFinallyBlock(ctx *FinallyBlockContext) {}

// ExitFinallyBlock is called when production finallyBlock is exited.
func (s *BaseapexListener) ExitFinallyBlock(ctx *FinallyBlockContext) {}

// EnterResourceSpecification is called when production resourceSpecification is entered.
func (s *BaseapexListener) EnterResourceSpecification(ctx *ResourceSpecificationContext) {}

// ExitResourceSpecification is called when production resourceSpecification is exited.
func (s *BaseapexListener) ExitResourceSpecification(ctx *ResourceSpecificationContext) {}

// EnterResources is called when production resources is entered.
func (s *BaseapexListener) EnterResources(ctx *ResourcesContext) {}

// ExitResources is called when production resources is exited.
func (s *BaseapexListener) ExitResources(ctx *ResourcesContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseapexListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseapexListener) ExitResource(ctx *ResourceContext) {}

// EnterForControl is called when production forControl is entered.
func (s *BaseapexListener) EnterForControl(ctx *ForControlContext) {}

// ExitForControl is called when production forControl is exited.
func (s *BaseapexListener) ExitForControl(ctx *ForControlContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BaseapexListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseapexListener) ExitForInit(ctx *ForInitContext) {}

// EnterEnhancedForControl is called when production enhancedForControl is entered.
func (s *BaseapexListener) EnterEnhancedForControl(ctx *EnhancedForControlContext) {}

// ExitEnhancedForControl is called when production enhancedForControl is exited.
func (s *BaseapexListener) ExitEnhancedForControl(ctx *EnhancedForControlContext) {}

// EnterForUpdate is called when production forUpdate is entered.
func (s *BaseapexListener) EnterForUpdate(ctx *ForUpdateContext) {}

// ExitForUpdate is called when production forUpdate is exited.
func (s *BaseapexListener) ExitForUpdate(ctx *ForUpdateContext) {}

// EnterParExpression is called when production parExpression is entered.
func (s *BaseapexListener) EnterParExpression(ctx *ParExpressionContext) {}

// ExitParExpression is called when production parExpression is exited.
func (s *BaseapexListener) ExitParExpression(ctx *ParExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseapexListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseapexListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterStatementExpression is called when production statementExpression is entered.
func (s *BaseapexListener) EnterStatementExpression(ctx *StatementExpressionContext) {}

// ExitStatementExpression is called when production statementExpression is exited.
func (s *BaseapexListener) ExitStatementExpression(ctx *StatementExpressionContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *BaseapexListener) EnterConstantExpression(ctx *ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *BaseapexListener) ExitConstantExpression(ctx *ConstantExpressionContext) {}

// EnterApexDbUpsertExpression is called when production apexDbUpsertExpression is entered.
func (s *BaseapexListener) EnterApexDbUpsertExpression(ctx *ApexDbUpsertExpressionContext) {}

// ExitApexDbUpsertExpression is called when production apexDbUpsertExpression is exited.
func (s *BaseapexListener) ExitApexDbUpsertExpression(ctx *ApexDbUpsertExpressionContext) {}

// EnterApexDbExpression is called when production apexDbExpression is entered.
func (s *BaseapexListener) EnterApexDbExpression(ctx *ApexDbExpressionContext) {}

// ExitApexDbExpression is called when production apexDbExpression is exited.
func (s *BaseapexListener) ExitApexDbExpression(ctx *ApexDbExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseapexListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseapexListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseapexListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseapexListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterCreator is called when production creator is entered.
func (s *BaseapexListener) EnterCreator(ctx *CreatorContext) {}

// ExitCreator is called when production creator is exited.
func (s *BaseapexListener) ExitCreator(ctx *CreatorContext) {}

// EnterCreatedName is called when production createdName is entered.
func (s *BaseapexListener) EnterCreatedName(ctx *CreatedNameContext) {}

// ExitCreatedName is called when production createdName is exited.
func (s *BaseapexListener) ExitCreatedName(ctx *CreatedNameContext) {}

// EnterInnerCreator is called when production innerCreator is entered.
func (s *BaseapexListener) EnterInnerCreator(ctx *InnerCreatorContext) {}

// ExitInnerCreator is called when production innerCreator is exited.
func (s *BaseapexListener) ExitInnerCreator(ctx *InnerCreatorContext) {}

// EnterArrayCreatorRest is called when production arrayCreatorRest is entered.
func (s *BaseapexListener) EnterArrayCreatorRest(ctx *ArrayCreatorRestContext) {}

// ExitArrayCreatorRest is called when production arrayCreatorRest is exited.
func (s *BaseapexListener) ExitArrayCreatorRest(ctx *ArrayCreatorRestContext) {}

// EnterMapCreatorRest is called when production mapCreatorRest is entered.
func (s *BaseapexListener) EnterMapCreatorRest(ctx *MapCreatorRestContext) {}

// ExitMapCreatorRest is called when production mapCreatorRest is exited.
func (s *BaseapexListener) ExitMapCreatorRest(ctx *MapCreatorRestContext) {}

// EnterSetCreatorRest is called when production setCreatorRest is entered.
func (s *BaseapexListener) EnterSetCreatorRest(ctx *SetCreatorRestContext) {}

// ExitSetCreatorRest is called when production setCreatorRest is exited.
func (s *BaseapexListener) ExitSetCreatorRest(ctx *SetCreatorRestContext) {}

// EnterClassCreatorRest is called when production classCreatorRest is entered.
func (s *BaseapexListener) EnterClassCreatorRest(ctx *ClassCreatorRestContext) {}

// ExitClassCreatorRest is called when production classCreatorRest is exited.
func (s *BaseapexListener) ExitClassCreatorRest(ctx *ClassCreatorRestContext) {}

// EnterExplicitGenericInvocation is called when production explicitGenericInvocation is entered.
func (s *BaseapexListener) EnterExplicitGenericInvocation(ctx *ExplicitGenericInvocationContext) {}

// ExitExplicitGenericInvocation is called when production explicitGenericInvocation is exited.
func (s *BaseapexListener) ExitExplicitGenericInvocation(ctx *ExplicitGenericInvocationContext) {}

// EnterNonWildcardTypeArguments is called when production nonWildcardTypeArguments is entered.
func (s *BaseapexListener) EnterNonWildcardTypeArguments(ctx *NonWildcardTypeArgumentsContext) {}

// ExitNonWildcardTypeArguments is called when production nonWildcardTypeArguments is exited.
func (s *BaseapexListener) ExitNonWildcardTypeArguments(ctx *NonWildcardTypeArgumentsContext) {}

// EnterTypeArgumentsOrDiamond is called when production typeArgumentsOrDiamond is entered.
func (s *BaseapexListener) EnterTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) {}

// ExitTypeArgumentsOrDiamond is called when production typeArgumentsOrDiamond is exited.
func (s *BaseapexListener) ExitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) {}

// EnterNonWildcardTypeArgumentsOrDiamond is called when production nonWildcardTypeArgumentsOrDiamond is entered.
func (s *BaseapexListener) EnterNonWildcardTypeArgumentsOrDiamond(ctx *NonWildcardTypeArgumentsOrDiamondContext) {
}

// ExitNonWildcardTypeArgumentsOrDiamond is called when production nonWildcardTypeArgumentsOrDiamond is exited.
func (s *BaseapexListener) ExitNonWildcardTypeArgumentsOrDiamond(ctx *NonWildcardTypeArgumentsOrDiamondContext) {
}

// EnterSuperSuffix is called when production superSuffix is entered.
func (s *BaseapexListener) EnterSuperSuffix(ctx *SuperSuffixContext) {}

// ExitSuperSuffix is called when production superSuffix is exited.
func (s *BaseapexListener) ExitSuperSuffix(ctx *SuperSuffixContext) {}

// EnterExplicitGenericInvocationSuffix is called when production explicitGenericInvocationSuffix is entered.
func (s *BaseapexListener) EnterExplicitGenericInvocationSuffix(ctx *ExplicitGenericInvocationSuffixContext) {
}

// ExitExplicitGenericInvocationSuffix is called when production explicitGenericInvocationSuffix is exited.
func (s *BaseapexListener) ExitExplicitGenericInvocationSuffix(ctx *ExplicitGenericInvocationSuffixContext) {
}

// EnterArguments is called when production arguments is entered.
func (s *BaseapexListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseapexListener) ExitArguments(ctx *ArgumentsContext) {}
