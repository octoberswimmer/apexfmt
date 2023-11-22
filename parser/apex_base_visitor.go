// Code generated from ./apex.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // apex
import "github.com/antlr4-go/antlr/v4"

type BaseapexVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseapexVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitImportDeclaration(ctx *ImportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitModifier(ctx *ModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitClassOrInterfaceModifier(ctx *ClassOrInterfaceModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitVariableModifier(ctx *VariableModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitTypeParameters(ctx *TypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitTypeBound(ctx *TypeBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitEnumConstants(ctx *EnumConstantsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitEnumConstant(ctx *EnumConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitClassBody(ctx *ClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitInterfaceBody(ctx *InterfaceBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitMemberDeclaration(ctx *MemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitMethodDeclaration(ctx *MethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitGenericMethodDeclaration(ctx *GenericMethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitGenericConstructorDeclaration(ctx *GenericConstructorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitPropertyDeclaration(ctx *PropertyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitPropertyBodyDeclaration(ctx *PropertyBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitConstDeclaration(ctx *ConstDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitConstantDeclarator(ctx *ConstantDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitGenericInterfaceMethodDeclaration(ctx *GenericInterfaceMethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitVariableInitializer(ctx *VariableInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitArrayInitializer(ctx *ArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitEnumConstantName(ctx *EnumConstantNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitType_(ctx *Type_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitTypeArguments(ctx *TypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitTypeArgument(ctx *TypeArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitQualifiedNameList(ctx *QualifiedNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitFormalParameters(ctx *FormalParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitFormalParameter(ctx *FormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitLastFormalParameter(ctx *LastFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitMethodBody(ctx *MethodBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitConstructorBody(ctx *ConstructorBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitAnnotationName(ctx *AnnotationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitElementValuePairs(ctx *ElementValuePairsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitElementValuePair(ctx *ElementValuePairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitElementValue(ctx *ElementValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitAnnotationTypeBody(ctx *AnnotationTypeBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitAnnotationTypeElementRest(ctx *AnnotationTypeElementRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitAnnotationMethodOrConstantRest(ctx *AnnotationMethodOrConstantRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitAnnotationMethodRest(ctx *AnnotationMethodRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitAnnotationConstantRest(ctx *AnnotationConstantRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitDefaultValue(ctx *DefaultValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitPropertyBlock(ctx *PropertyBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitGetter(ctx *GetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitSetter(ctx *SetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitCatchClause(ctx *CatchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitCatchType(ctx *CatchTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitFinallyBlock(ctx *FinallyBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitResourceSpecification(ctx *ResourceSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitResources(ctx *ResourcesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitResource(ctx *ResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitForControl(ctx *ForControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitForInit(ctx *ForInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitEnhancedForControl(ctx *EnhancedForControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitForUpdate(ctx *ForUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitParExpression(ctx *ParExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitStatementExpression(ctx *StatementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitConstantExpression(ctx *ConstantExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitApexDbUpsertExpression(ctx *ApexDbUpsertExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitApexDbExpression(ctx *ApexDbExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitCreator(ctx *CreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitCreatedName(ctx *CreatedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitInnerCreator(ctx *InnerCreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitArrayCreatorRest(ctx *ArrayCreatorRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitMapCreatorRest(ctx *MapCreatorRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitSetCreatorRest(ctx *SetCreatorRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitClassCreatorRest(ctx *ClassCreatorRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitExplicitGenericInvocation(ctx *ExplicitGenericInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitNonWildcardTypeArguments(ctx *NonWildcardTypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitNonWildcardTypeArgumentsOrDiamond(ctx *NonWildcardTypeArgumentsOrDiamondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitSuperSuffix(ctx *SuperSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitExplicitGenericInvocationSuffix(ctx *ExplicitGenericInvocationSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseapexVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}
