// Code generated from ./apex.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // apex
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by apexParser.
type apexVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by apexParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by apexParser#packageDeclaration.
	VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#importDeclaration.
	VisitImportDeclaration(ctx *ImportDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#modifier.
	VisitModifier(ctx *ModifierContext) interface{}

	// Visit a parse tree produced by apexParser#classOrInterfaceModifier.
	VisitClassOrInterfaceModifier(ctx *ClassOrInterfaceModifierContext) interface{}

	// Visit a parse tree produced by apexParser#variableModifier.
	VisitVariableModifier(ctx *VariableModifierContext) interface{}

	// Visit a parse tree produced by apexParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#typeParameters.
	VisitTypeParameters(ctx *TypeParametersContext) interface{}

	// Visit a parse tree produced by apexParser#typeParameter.
	VisitTypeParameter(ctx *TypeParameterContext) interface{}

	// Visit a parse tree produced by apexParser#typeBound.
	VisitTypeBound(ctx *TypeBoundContext) interface{}

	// Visit a parse tree produced by apexParser#enumDeclaration.
	VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#enumConstants.
	VisitEnumConstants(ctx *EnumConstantsContext) interface{}

	// Visit a parse tree produced by apexParser#enumConstant.
	VisitEnumConstant(ctx *EnumConstantContext) interface{}

	// Visit a parse tree produced by apexParser#enumBodyDeclarations.
	VisitEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) interface{}

	// Visit a parse tree produced by apexParser#interfaceDeclaration.
	VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by apexParser#classBody.
	VisitClassBody(ctx *ClassBodyContext) interface{}

	// Visit a parse tree produced by apexParser#interfaceBody.
	VisitInterfaceBody(ctx *InterfaceBodyContext) interface{}

	// Visit a parse tree produced by apexParser#classBodyDeclaration.
	VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#memberDeclaration.
	VisitMemberDeclaration(ctx *MemberDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#methodDeclaration.
	VisitMethodDeclaration(ctx *MethodDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#genericMethodDeclaration.
	VisitGenericMethodDeclaration(ctx *GenericMethodDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#constructorDeclaration.
	VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#genericConstructorDeclaration.
	VisitGenericConstructorDeclaration(ctx *GenericConstructorDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#fieldDeclaration.
	VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#propertyDeclaration.
	VisitPropertyDeclaration(ctx *PropertyDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#propertyBodyDeclaration.
	VisitPropertyBodyDeclaration(ctx *PropertyBodyDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#interfaceBodyDeclaration.
	VisitInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#interfaceMemberDeclaration.
	VisitInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#constDeclaration.
	VisitConstDeclaration(ctx *ConstDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#constantDeclarator.
	VisitConstantDeclarator(ctx *ConstantDeclaratorContext) interface{}

	// Visit a parse tree produced by apexParser#interfaceMethodDeclaration.
	VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#genericInterfaceMethodDeclaration.
	VisitGenericInterfaceMethodDeclaration(ctx *GenericInterfaceMethodDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#variableDeclarators.
	VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{}

	// Visit a parse tree produced by apexParser#variableDeclarator.
	VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{}

	// Visit a parse tree produced by apexParser#variableDeclaratorId.
	VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{}

	// Visit a parse tree produced by apexParser#variableInitializer.
	VisitVariableInitializer(ctx *VariableInitializerContext) interface{}

	// Visit a parse tree produced by apexParser#arrayInitializer.
	VisitArrayInitializer(ctx *ArrayInitializerContext) interface{}

	// Visit a parse tree produced by apexParser#enumConstantName.
	VisitEnumConstantName(ctx *EnumConstantNameContext) interface{}

	// Visit a parse tree produced by apexParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by apexParser#classOrInterfaceType.
	VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by apexParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by apexParser#typeArguments.
	VisitTypeArguments(ctx *TypeArgumentsContext) interface{}

	// Visit a parse tree produced by apexParser#typeArgument.
	VisitTypeArgument(ctx *TypeArgumentContext) interface{}

	// Visit a parse tree produced by apexParser#qualifiedNameList.
	VisitQualifiedNameList(ctx *QualifiedNameListContext) interface{}

	// Visit a parse tree produced by apexParser#formalParameters.
	VisitFormalParameters(ctx *FormalParametersContext) interface{}

	// Visit a parse tree produced by apexParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by apexParser#formalParameter.
	VisitFormalParameter(ctx *FormalParameterContext) interface{}

	// Visit a parse tree produced by apexParser#lastFormalParameter.
	VisitLastFormalParameter(ctx *LastFormalParameterContext) interface{}

	// Visit a parse tree produced by apexParser#methodBody.
	VisitMethodBody(ctx *MethodBodyContext) interface{}

	// Visit a parse tree produced by apexParser#constructorBody.
	VisitConstructorBody(ctx *ConstructorBodyContext) interface{}

	// Visit a parse tree produced by apexParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}

	// Visit a parse tree produced by apexParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by apexParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by apexParser#annotationName.
	VisitAnnotationName(ctx *AnnotationNameContext) interface{}

	// Visit a parse tree produced by apexParser#elementValuePairs.
	VisitElementValuePairs(ctx *ElementValuePairsContext) interface{}

	// Visit a parse tree produced by apexParser#elementValuePair.
	VisitElementValuePair(ctx *ElementValuePairContext) interface{}

	// Visit a parse tree produced by apexParser#elementValue.
	VisitElementValue(ctx *ElementValueContext) interface{}

	// Visit a parse tree produced by apexParser#elementValueArrayInitializer.
	VisitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) interface{}

	// Visit a parse tree produced by apexParser#annotationTypeDeclaration.
	VisitAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#annotationTypeBody.
	VisitAnnotationTypeBody(ctx *AnnotationTypeBodyContext) interface{}

	// Visit a parse tree produced by apexParser#annotationTypeElementDeclaration.
	VisitAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#annotationTypeElementRest.
	VisitAnnotationTypeElementRest(ctx *AnnotationTypeElementRestContext) interface{}

	// Visit a parse tree produced by apexParser#annotationMethodOrConstantRest.
	VisitAnnotationMethodOrConstantRest(ctx *AnnotationMethodOrConstantRestContext) interface{}

	// Visit a parse tree produced by apexParser#annotationMethodRest.
	VisitAnnotationMethodRest(ctx *AnnotationMethodRestContext) interface{}

	// Visit a parse tree produced by apexParser#annotationConstantRest.
	VisitAnnotationConstantRest(ctx *AnnotationConstantRestContext) interface{}

	// Visit a parse tree produced by apexParser#defaultValue.
	VisitDefaultValue(ctx *DefaultValueContext) interface{}

	// Visit a parse tree produced by apexParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by apexParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by apexParser#localVariableDeclarationStatement.
	VisitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) interface{}

	// Visit a parse tree produced by apexParser#localVariableDeclaration.
	VisitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by apexParser#propertyBlock.
	VisitPropertyBlock(ctx *PropertyBlockContext) interface{}

	// Visit a parse tree produced by apexParser#getter.
	VisitGetter(ctx *GetterContext) interface{}

	// Visit a parse tree produced by apexParser#setter.
	VisitSetter(ctx *SetterContext) interface{}

	// Visit a parse tree produced by apexParser#catchClause.
	VisitCatchClause(ctx *CatchClauseContext) interface{}

	// Visit a parse tree produced by apexParser#catchType.
	VisitCatchType(ctx *CatchTypeContext) interface{}

	// Visit a parse tree produced by apexParser#finallyBlock.
	VisitFinallyBlock(ctx *FinallyBlockContext) interface{}

	// Visit a parse tree produced by apexParser#resourceSpecification.
	VisitResourceSpecification(ctx *ResourceSpecificationContext) interface{}

	// Visit a parse tree produced by apexParser#resources.
	VisitResources(ctx *ResourcesContext) interface{}

	// Visit a parse tree produced by apexParser#resource.
	VisitResource(ctx *ResourceContext) interface{}

	// Visit a parse tree produced by apexParser#forControl.
	VisitForControl(ctx *ForControlContext) interface{}

	// Visit a parse tree produced by apexParser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by apexParser#enhancedForControl.
	VisitEnhancedForControl(ctx *EnhancedForControlContext) interface{}

	// Visit a parse tree produced by apexParser#forUpdate.
	VisitForUpdate(ctx *ForUpdateContext) interface{}

	// Visit a parse tree produced by apexParser#parExpression.
	VisitParExpression(ctx *ParExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by apexParser#statementExpression.
	VisitStatementExpression(ctx *StatementExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#constantExpression.
	VisitConstantExpression(ctx *ConstantExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#apexDbUpsertExpression.
	VisitApexDbUpsertExpression(ctx *ApexDbUpsertExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#apexDbExpression.
	VisitApexDbExpression(ctx *ApexDbExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by apexParser#creator.
	VisitCreator(ctx *CreatorContext) interface{}

	// Visit a parse tree produced by apexParser#createdName.
	VisitCreatedName(ctx *CreatedNameContext) interface{}

	// Visit a parse tree produced by apexParser#innerCreator.
	VisitInnerCreator(ctx *InnerCreatorContext) interface{}

	// Visit a parse tree produced by apexParser#arrayCreatorRest.
	VisitArrayCreatorRest(ctx *ArrayCreatorRestContext) interface{}

	// Visit a parse tree produced by apexParser#mapCreatorRest.
	VisitMapCreatorRest(ctx *MapCreatorRestContext) interface{}

	// Visit a parse tree produced by apexParser#setCreatorRest.
	VisitSetCreatorRest(ctx *SetCreatorRestContext) interface{}

	// Visit a parse tree produced by apexParser#classCreatorRest.
	VisitClassCreatorRest(ctx *ClassCreatorRestContext) interface{}

	// Visit a parse tree produced by apexParser#explicitGenericInvocation.
	VisitExplicitGenericInvocation(ctx *ExplicitGenericInvocationContext) interface{}

	// Visit a parse tree produced by apexParser#nonWildcardTypeArguments.
	VisitNonWildcardTypeArguments(ctx *NonWildcardTypeArgumentsContext) interface{}

	// Visit a parse tree produced by apexParser#typeArgumentsOrDiamond.
	VisitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) interface{}

	// Visit a parse tree produced by apexParser#nonWildcardTypeArgumentsOrDiamond.
	VisitNonWildcardTypeArgumentsOrDiamond(ctx *NonWildcardTypeArgumentsOrDiamondContext) interface{}

	// Visit a parse tree produced by apexParser#superSuffix.
	VisitSuperSuffix(ctx *SuperSuffixContext) interface{}

	// Visit a parse tree produced by apexParser#explicitGenericInvocationSuffix.
	VisitExplicitGenericInvocationSuffix(ctx *ExplicitGenericInvocationSuffixContext) interface{}

	// Visit a parse tree produced by apexParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}
}
